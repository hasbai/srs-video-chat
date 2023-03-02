export class SrsError extends Error {
  constructor(
      public name: string,
      public message: string,
      public stack = new Error().stack
  ) {
    super()
  }

  toString() {
    return `${this.name}: ${this.message}`
  }
}

// Depends on adapter-7.4.0.min.js from https://github.com/webrtc/adapter
// Async-await-promise based SRS RTC Player.
export class SrsRtcPlayer {
  private pc = new RTCPeerConnection()
  public stream = new MediaStream()
  public audioEnabled = true
  public videoEnabled = true

  constructor() {
    this.pc.ontrack = (event) => {
      // https://webrtc.org/getting-started/remote-streams
      this.stream.addTrack(event.track)
    }
  }

  async play(url: string) {
    this.pc.addTransceiver('audio', {direction: 'recvonly'})
    this.pc.addTransceiver('video', {direction: 'recvonly'})

    return await this.negotiate(url, 'play')
  }

  async publish(url: string) {
    this.pc.addTransceiver('audio', {direction: 'sendonly'})
    this.pc.addTransceiver('video', {direction: 'sendonly'})

    if (
        !navigator.mediaDevices &&
        window.location.protocol === 'http:' &&
        window.location.hostname !== 'localhost'
    ) {
      throw new SrsError(
          'Https Required',
          `Please use HTTPS or localhost to publish, read https://github.com/ossrs/srs/issues/2762#issuecomment-983147576`
      )
    }

    const stream = await navigator.mediaDevices.getUserMedia({
      audio: true,
      video: true,
    })

    // @see https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/addStream#Migrating_to_addTrack
    stream.getTracks().forEach((track) => {
      let sender = this.pc.addTrack(track)
      this.stream.addTrack(track)
    })

    return await this.negotiate(url, 'publish')
  }

  async negotiate(url: string, type: string) {
    const offer = await this.pc.createOffer()
    await this.pc.setLocalDescription(offer)
    const data = {
      streamurl: url,
      sdp: offer.sdp,
    }
    const api = `/rtc/v1/${type}/`
    const r = await fetch(api, {
      method: 'POST', // or 'PUT'
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
    const session = (await r.json()) as any

    await this.pc.setRemoteDescription(
        new RTCSessionDescription({type: 'answer', sdp: session.sdp})
    )

    return session
  }

  changeMediaStatus(type: string) {
    switch (type) {
      case 'audio':
        this.audioEnabled = !this.audioEnabled
        this.stream.getAudioTracks().forEach((track: MediaStreamTrack) => {
          track.enabled = this.audioEnabled
        })
        break
      case 'video':
        this.videoEnabled = !this.videoEnabled
        this.stream.getVideoTracks().forEach((track: MediaStreamTrack) => {
          track.enabled = this.videoEnabled
        })
    }
  }

  // Close the publisher.
  close() {
    this.pc.close()
    this.pc = new RTCPeerConnection()
    this.stream = new MediaStream()
  }
}
