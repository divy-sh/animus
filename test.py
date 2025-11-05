import argparse, socket, threading, time, binascii
from datetime import datetime

def hexdump(data):
    hexs = binascii.hexlify(data).decode('ascii')
    # group into bytes
    pairs = [hexs[i:i+2] for i in range(0, len(hexs), 2)]
    ascii_repr = ''.join((chr(int(b,16)) if 32 <= int(b,16) < 127 else '.') for b in pairs)
    return ' '.join(pairs), ascii_repr

def relay(src, dst, label):
    try:
        while True:
            data = src.recv(4096)
            if not data:
                break
            t = datetime.now().isoformat(timespec='milliseconds')
            hexs, ascii_repr = hexdump(data)
            print(f"[{t}] -> {label} ({len(data)} bytes)\nHEX: {hexs}\nASCII: {ascii_repr}\n")
            dst.sendall(data)
    except Exception as e:
        print(f"[!] relay {label} error: {e}")
    finally:
        try: dst.shutdown(socket.SHUT_WR)
        except: pass

def handle_client(client_sock, remote_host, remote_port, addr):
    try:
        server_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server_sock.connect((remote_host, remote_port))
    except Exception as e:
        print(f"[!] Cannot connect to remote {remote_host}:{remote_port}: {e}")
        client_sock.close()
        return

    # start relay threads
    t1 = threading.Thread(target=relay, args=(client_sock, server_sock, f"client->{remote_host}:{remote_port}"), daemon=True)
    t2 = threading.Thread(target=relay, args=(server_sock, client_sock, f"{remote_host}:{remote_port}->client"), daemon=True)
    t1.start(); t2.start()
    # wait until both finish
    t1.join(); t2.join()
    client_sock.close(); server_sock.close()
    print(f"[{datetime.now().isoformat(timespec='seconds')}] Connection from {addr} closed\n")

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--listen", default="127.0.0.1")
    parser.add_argument("--lport", type=int, default=63790)
    parser.add_argument("--remote", default="127.0.0.1")
    parser.add_argument("--rport", type=int, default=6379)
    args = parser.parse_args()

    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    sock.bind((args.listen, args.lport))
    sock.listen(16)
    print(f"Listening on {args.listen}:{args.lport}, forwarding to {args.remote}:{args.rport}")
    try:
        while True:
            cl, addr = sock.accept()
            print(f"[{datetime.now().isoformat(timespec='seconds')}] Accepted connection from {addr}")
            threading.Thread(target=handle_client, args=(cl, args.remote, args.rport, addr), daemon=True).start()
    except KeyboardInterrupt:
        print("Shutting down.")
    finally:
        sock.close()

if __name__ == "__main__":
    main()
