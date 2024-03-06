"""The Python implementation of the gRPC route guide client."""

from __future__ import print_function
import time
import logging
import random

import grpc
import tts_pb2_grpc
import tts_pb2

# import playsound


with grpc.insecure_channel("localhost:50052") as channel:
    stub = tts_pb2_grpc.TextToSpeechStub(channel)
    print("Getting the full made voice")
    start = time.time()
    response = stub.Synthesize(
        tts_pb2.SynthesizeRequest(
            text="XTTS builds upon the recent advancements in autoregressive models, such as Tortoise, Vall-E, and Soundstorm, which are based on language models trained on discrete audio representations. XTTS utilizes a VQ-VAE model to discretize the audio into audio tokens. Subsequently, it employs a GPT model to predict these audio tokens based on the input text and speaker latents. The speaker latents are computed by a stack of self-attention layers. The output of the GPT model is passed on to a decoder model that outputs the audio signal. We employ the Tortoise methodology for XTTS-v1, which combines a diffusion model and UnivNet vocoder.",
            language_code="en",
        )
    )
    # save the audio to a file
    # print the response code
    # print(f"Response code: {response.code}")
    # print(response.audio)
    # print("Audio saved to output.wav")
    print("Response for inference Full received - time taken: ", time.time() - start)
    print("Stream the audio")
    response_iterator = stub.SynthesizeStream(
        tts_pb2.SynthesizeRequest(
            text="XTTS builds upon the recent advancements in autoregressive models, such as Tortoise, Vall-E, and Soundstorm, which are based on language models trained on discrete audio representations. XTTS utilizes a VQ-VAE model to discretize the audio into audio tokens. Subsequently, it employs a GPT model to predict these audio tokens based on the input text and speaker latents. The speaker latents are computed by a stack of self-attention layers. The output of the GPT model is passed on to a decoder model that outputs the audio signal. We employ the Tortoise methodology for XTTS-v1, which combines a diffusion model and UnivNet vocoder. This approach involves using the diffusion model to transform the GPT outputs into spectrogram frames, and then utilizing UnivNet to generate the ultimate audio signal.",
            language_code="en",
        )
    )
    i = 0
    start = time.time()
    for response in response_iterator:
        # play the audio to speaker
        # print("Playing the audio")
        # better to save it
        now = time.time()
        print(f"Response {i} received - time: {time.time() - start}")
    print("Response for inference Stream received - time taken: ", time.time() - start)

