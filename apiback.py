from flask import Flask, request, Response
import os
from flask_cors import CORS  # Importa la extensión Flask-CORS

app = Flask(__name__)
CORS(app)



# Rutas a los archivos generados por el generador DASH
VIDEO_PATH = 'video.mp4'
MPD_PATH = 'cachos/dash_out.mpd'
SEGMENTS_PATH = 'cachos/'

@app.route('/video_stream.mpd')
def video_stream():
    with open(MPD_PATH, 'rb') as mpd_file:
        mpd_data = mpd_file.read()
    return Response(mpd_data, content_type='application/dash+xml')

@app.route('/init-stream<stream_id>.m4s')
def init_segment(stream_id):
    init_segment_path = os.path.join(SEGMENTS_PATH, f'init-stream{stream_id}.m4s')
    
    if os.path.exists(init_segment_path):
        with open(init_segment_path, 'rb') as init_segment_file:
            data = init_segment_file.read()
            return Response(data, content_type='video/mp4')
    else:
        return Response("Initialization segment not found", status=404)

@app.route('/chunk-stream<stream_id>-<segment_number>.m4s')
def video_segment(stream_id, segment_number):
    segment_filename = f'chunk-stream{stream_id}-{segment_number.zfill(5)}.m4s'
    segment_path = os.path.join(SEGMENTS_PATH, segment_filename)
    
    if os.path.exists(segment_path):
        with open(segment_path, 'rb') as segment_file:
            data = segment_file.read()
            return Response(data, content_type='video/mp4')
    else:
        return Response("Segment not found", status=404)

@app.route('/')
def index():
    return "DASH Video Streaming API"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080, debug=True)
