export class SrsRtcSignaling {
  private ws: WebSocket

  // The schema is ws or wss, host is ip or ip:port, display is nickname
  // of user to join the room.
  connect(room: string, user: string) {
    const url = window.location.origin.replace('http', 'ws') + '/ws'
    this.ws = new WebSocket(url + '?room=' + room + '&user=' + user)

    this.ws.onmessage = (event) => {
      const r = JSON.parse(event.data)
      console.log(r)
    }

    return new Promise((resolve, reject) => {
      this.ws.onopen = function (event) {
        resolve(event)
      }

      this.ws.onerror = function (event) {
        reject(event)
      }
    })
  }

  send(message: string) {
    this.ws.send(message)
  }

  close() {
    this.ws && this.ws.close()
  }
}

// Parse params in query string.
function SrsRtcSignalingParse(location: Location) {
  let query: string | null = location.href.split('?')[1]
  query = query ? '?' + query : null

  let wsSchema = location.href.split('wss=')[1]
  wsSchema = wsSchema
    ? wsSchema.split('&')[0]
    : location.protocol === 'http:'
    ? 'ws'
    : 'wss'

  let wsHost = location.href.split('wsh=')[1]
  wsHost = wsHost ? wsHost.split('&')[0] : location.hostname

  let wsPort = location.href.split('wsp=')[1]
  wsPort = wsPort ? wsPort.split('&')[0] : location.host.split(':')[1]

  let host = location.href.split('host=')[1]
  host = host ? host.split('&')[0] : location.hostname

  let room: string | null = location.href.split('room=')[1]
  room = room ? room.split('&')[0] : null

  let display = location.href.split('display=')[1]
  display = display
    ? display.split('&')[0]
    : (new Date().getTime() * Math.random() * 100).toString(16).slice(0, 7)

  let autostart: string = location.href.split('autostart=')[1]
  const isAutoStart = autostart && autostart.split('&')[0] === 'true'

  // Remove data in query.
  let rawQuery = query
  if (query) {
    query = query.replace('wss=' + wsSchema, '')
    query = query.replace('wsh=' + wsHost, '')
    query = query.replace('wsp=' + wsPort, '')
    query = query.replace('host=' + host, '')
    if (room) {
      query = query.replace('room=' + room, '')
    }
    query = query.replace('display=' + display, '')
    query = query.replace('autostart=' + autostart, '')

    while (query.indexOf('&&') >= 0) {
      query = query.replace('&&', '&')
    }
    query = query.replace('?&', '?')
    if (query.lastIndexOf('?') === query.length - 1) {
      query = query.slice(0, query.length - 1)
    }
    if (query.lastIndexOf('&') === query.length - 1) {
      query = query.slice(0, query.length - 1)
    }
  }

  // Regenerate the host of websocket.
  wsHost = wsPort ? wsHost.split(':')[0] + ':' + wsPort : wsHost

  return {
    query: query,
    rawQuery: rawQuery,
    wsSchema: wsSchema,
    wsHost: wsHost,
    host: host,
    room: room,
    display: display,
    autostart: isAutoStart,
  }
}
