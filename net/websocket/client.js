const wsc = new WebSocket(
  "wss://localhost:8001/socketserver",
);

const exampleSocket = new WebSocket(
  "ws://localhost:8001/socketserver",
  // ["webpack-hmr"]
);

exampleSocket.onmessage = (event) => {
  console.log(event.data);
};
exampleSocket.onopen = (event) => {
  const msg = {
    type: "message",
    date: Date.now(),
  };

  // Send the msg object as a JSON-formatted string.
  exampleSocket.send(JSON.stringify(msg));
};
