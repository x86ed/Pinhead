const signup = (name, initials) => {
  const myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");

  const raw = JSON.stringify({
    name: name,
    password: name + initials,
    role: "user",
    initials: initials,
  });

  const requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("/signup", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.log("error", error));
};

const signin = (name, initials) => {
  const myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");

  const raw = JSON.stringify({
    name: name,
    password: name + initials,
  });

  const requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("http://localhost:8080/signin", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.log("error", error));
};

var myHeaders = new Headers();
myHeaders.append(
  "Authorization",
  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTYwMjExNDIsIm5hbWUiOiJtYXR0aGV3Iiwicm9sZSI6InVzZXIifQ.E8SRBEwfxMGnfGvjs5lYGpvtT8nVXxXlqJIFhKHzcDQ"
);

var requestOptions = {
  method: "GET",
  headers: myHeaders,
  redirect: "follow",
};

fetch("http://localhost:8080/user", requestOptions)
  .then((response) => response.text())
  .then((result) => console.log(result))
  .catch((error) => console.log("error", error));

const getSocketURI = () => {
  const myHeaders = new Headers();
  myHeaders.append(
    "Authorization",
    "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTYwMjExNDIsIm5hbWUiOiJtYXR0aGV3Iiwicm9sZSI6InVzZXIifQ.E8SRBEwfxMGnfGvjs5lYGpvtT8nVXxXlqJIFhKHzcDQ"
  );

  const requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: null,
    redirect: "follow",
  };

  fetch("http://localhost:8080/wsurl", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.log("error", error));
};

var output = document.getElementById("output");
var input = document.getElementById("input");
let ws;

var print = function (message) {
  console.log(message);
};

const sendS = () => {
  if (!ws) {
    return false;
  }
  print("SEND: S");
  ws.send("S");
  return false;
};

const sendL = () => {
  if (!ws) {
    return false;
  }
  print("SEND: L");
  ws.send("L");
  return false;
};

const sendLU = () => {
  if (!ws) {
    return false;
  }
  print("SEND: LU");
  ws.send("LU");
  return false;
};

const sendLD = () => {
  if (!ws) {
    return false;
  }
  print("SEND: LD");
  ws.send("LD");
  return false;
};

const sendRU = () => {
  if (!ws) {
    return false;
  }
  print("SEND: RU");
  ws.send("RU");
  return false;
};

const sendRD = () => {
  if (!ws) {
    return false;
  }
  print("SEND: RD");
  ws.send("RD");
  return false;
};

document.querySelector("body").onload = function (evt) {
  if (ws) {
    return false;
  }
  ws = new WebSocket(getSocketURI());

  ws.onopen = function (evt) {
    print("OPEN");
  };
  ws.onclose = function (evt) {
    print("CLOSE");
    ws = null;
  };
  ws.onmessage = function (evt) {
    print("RESPONSE: " + evt.data);
  };
  ws.onerror = function (evt) {
    print("ERROR: " + evt.data);
  };

  document.getElementById("startbutton").onclick = sendS;

  document.getElementById("launchbutton").onclick = sendL;

  document.getElementById("leftbutton").onmousedown = sendLU;

  document.getElementById("rightbutton").onmousedown = sendRU;

  document.getElementById("leftbutton").onmouseup = sendLD;

  document.getElementById("rightbutton").onmouseup = sendRD;
  return false;
};

document.getElementById("send").onclick = function () {
  if (!ws) {
    return false;
  }
  print("SEND: " + input.value);
  ws.send(input.value);
  return false;
};

document.getElementById("close").onclick = function () {
  if (!ws) {
    return false;
  }
  ws.close();
  return false;
};
