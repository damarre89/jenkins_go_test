from http.server import BaseHTTPRequestHandler, HTTPServer

class Handler(BaseHTTPRequestHandler):
    def do_POST(self):
        length = int(self.headers.get("Content-Length", 0))
        body = self.rfile.read(length)

        print("\n=== REQUEST ===")
        print(f"Method: {self.command}")
        print(f"Path:   {self.path}")
        print("\nHeaders:")
        for name, value in self.headers.items():
            print(f"{name}: {value}")

        print("\nBody:")
        print(body.decode("utf-8", errors="replace"))
        print("==============\n")

        self.send_response(200)
        self.end_headers()
        self.wfile.write(b"OK")

    def log_message(self, format, *args):
        pass

HTTPServer(("0.0.0.0", 9090), Handler).serve_forever().
