# For CogLog (NatHacks2023)
# Load BCI-decoder model and run web-server
# From https://github.com/fwillett/speechBCI

baseDir = '.'
import os
from glob import glob
from pathlib import Path
os.environ["CUDA_DEVICE_ORDER"]="PCI_BUS_ID"
os.environ["CUDA_VISIBLE_DEVICES"]=""

import numpy as np
from omegaconf import OmegaConf
import tensorflow as tf
from neuralDecoder.neuralSequenceDecoder import NeuralSequenceDecoder
import neuralDecoder.utils.lmDecoderUtils as lmDecoderUtils

print("LANGUAGE MODEL LOADING...")

#loads the language model, could take a while and requires ~60 GB of memory
lmDir = baseDir+'/languageModel'
ngramDecoder = lmDecoderUtils.build_lm_decoder(
    lmDir,
    acoustic_scale=0.8, #1.2
    nbest=1,
    beam=18
)

print("CONFIGURING...")

#evaluate the RNN on the test partition and competitionHoldOut partition
testDirs = ['test','competitionHoldOut']
trueTranscriptions = [[]]
decodedTranscriptions = [[]]

ckptDir = baseDir + '/derived/rnns/baselineRelease'

args = OmegaConf.load(os.path.join(ckptDir, 'args.yaml'))
args['loadDir'] = ckptDir
args['mode'] = 'infer'
args['loadCheckpointIdx'] = None

for x in range(len(args['dataset']['datasetProbabilityVal'])):
    args['dataset']['datasetProbabilityVal'][x] = 0.0

for sessIdx in range(1):
    args['dataset']['datasetProbabilityVal'][sessIdx] = 1.0
    args['dataset']['dataDir'][sessIdx] = baseDir+'/derived/tfRecords'
args['testDir'] = testDirs[0]

print("INITIALIZING RNN...")

# Initialize model
tf.compat.v1.reset_default_graph()
nsd = NeuralSequenceDecoder(args)

print("INFERENCING...")

# Inference
out = nsd.inference()
decoder_out = lmDecoderUtils.cer_with_lm_decoder(ngramDecoder, out, outputType='speech_sil', blankPenalty=np.log(2))

print("DONE")

def _ascii_to_text(text):
    endIdx = np.argwhere(text==0)
    return ''.join([chr(char) for char in text[0:endIdx[0,0]]])

for x in range(out['transcriptions'].shape[0]):
    trueTranscriptions[0].append(_ascii_to_text(out['transcriptions'][x,:]))
decodedTranscriptions[0] = decoder_out['decoded_transcripts']

from neuralDecoder.utils.lmDecoderUtils import _cer_and_wer as cer_and_wer

#get word error rate and phoneme error rate for the test set (cer is actually phoneme error rate here)
cer, wer = cer_and_wer(decodedTranscriptions[0], trueTranscriptions[0], outputType='speech_sil', returnCI=True)

#print word error rate
print(wer)

#print the sentence predictions for the test set
print(decodedTranscriptions[0])
print(trueTranscriptions[0])

# server
from http.server import BaseHTTPRequestHandler, HTTPServer
import time

hostName = "10.128.0.2"
serverPort = 8080

class MyServer(BaseHTTPRequestHandler):
    def do_POST(self):
        content_len = int(self.headers.get('Content-Length'))
        post_body = self.rfile.read(content_len)

        num = int(post_body.decode('utf-8'))
        print(num)

        if num < 0 or num > len(decodedTranscriptions[0]):
            self.send_response(400)
            self.send_header("Content-type", "application/json")
            self.end_headers()
            self.wfile.write(bytes('{"err": "invalid index"}', "utf-8"))
        else:
            self.send_response(200)
            self.send_header("Content-type", "application/json")
            self.end_headers()
            self.wfile.write(bytes('{"text": "'+decodedTranscriptions[0][num]+'"}', "utf-8"))

webServer = HTTPServer((hostName, serverPort), MyServer)
print("Server started http://%s:%s" % (hostName, serverPort))

try:
    webServer.serve_forever()
except KeyboardInterrupt:
    pass

webServer.server_close()
print("Server stopped.")
