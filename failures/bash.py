import subprocess
import time
import BaseHTTPServer

def run(server_class=HTTPServer, handler_class=BaseHTTPRequestHandler):
    server_address = ('', 8000)
    httpd = server_class(server_address, handler_class)
    httpd.serve_forever()




subprocess.check_call(['/my/file/path/programname.sh', 'arg1', 'arg2', arg3])
