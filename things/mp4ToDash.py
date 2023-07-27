# import required libraries
from vidgear.gears import StreamGear
import os,sys

print("\n\n\nESTAMO ACTIVO PAPI\n\n\n")
source = sys.argv[1]
id = sys.argv[2]
title = sys.argv[3]

def video(source,id):
# activate Single-Source Mode and also define various streams
    stream_params = {
        "-video_source": f"backend/pb_data/storage/4ronlqa5jkr2oda/{id}/{source}",
        "-streams": [
            {"-resolution": "1920x1080", "-video_bitrate": "4000k"},  # Stream1: 1920x1080 at 4000kbs bitrate
            {"-resolution": "1280x720", "-framerate": 30.0},  # Stream2: 1280x720 at 30fps framerate
            {"-resolution": "640x360", "-framerate": 60.0},  # Stream3: 640x360 at 60fps framerate
            {"-resolution": "320x240", "-video_bitrate": "500k"},  # Stream3: 320x240 at 500kbs bitrate
        ],
    }
    return stream_params


stream_params = video(source, id)
# main_directory = "../cachos"

# # Si el directorio principal no existe, créalo
# if not os.path.exists(main_directory):
#     os.makedirs(main_directory)
print(f"backend/pb_data/storage/4ronlqa5jkr2oda/{id}/dash_out.mpd")
# # Luego, crea el subdirectorio con el nombre `title`
try:
    #os.mkdir(f"{main_directory}/{title}")
    print("transformandooo...")
     # describe a suitable manifest-file location/name and assign params
    out = os.path.join(os.getcwd(), "backend", "pb_data", "storage", "4ronlqa5jkr2oda", id, "dash_out.mpd")
    streamer = StreamGear(output=os.path.join(os.getcwd(), "backend", "pb_data", "storage", "4ronlqa5jkr2oda", id, "dash_out.mpd"), **stream_params)
    # trancode source
    streamer.transcode_source()
    # terminate
    streamer.terminate()
    print("hechooo")
except Exception as err:
    print(err)
