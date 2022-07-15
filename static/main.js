const signup = (name,initials) =>{
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    const raw = JSON.stringify({
        "name": name,
        "password": name+initials,
        "role": "user",
        "initials": initials
    });
    
    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };
    
    fetch("/signup", requestOptions)
        .then(response => response.text())
        .then(result => updateList(result))
        .catch(error => console.log('error', error));
}

const updateList = (result) => {
    
    let queueList = document.getElementById('queue-list');
    
    let arr = [];
    for (const element of JSON.parse(result)) {
        let listItem = document.createElement('li');
        listItem.textContent = `${element.Name} - ${element.Initials}`;
        arr.push(listItem);
    }
    
    arr[0].classList.add("user");
    
    queueList.replaceChildren(...arr);
}

const signin = (name,initials)=>{
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    const raw = JSON.stringify({
        "name": name,
        "password": name+initials
    });

    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    fetch("http://localhost:8080/signin", requestOptions)
    .then(response => response.text())
    .then(result => console.log(result))
    .catch(error => console.log('error', error));
}

const myHeaders = new Headers();
myHeaders.append("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTU2Njg5MjAsIm5hbWUiOiJkZXJwIiwicm9sZSI6InVzZXIifQ.w09Rwoa7X0Fu5aMXrvQ5KMwA5VeSpMhSJ1j24snTdJU");

const requestOptions = {
  method: 'GET',
  headers: myHeaders,
  redirect: 'follow'
};

fetch("http://localhost:8080/user", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));

const getSocketURI = () => {
    let loc = window.location, new_uri;
    if (loc.protocol === "https:") {
        new_uri = "wss:";
    } else {
        new_uri = "ws:";
    }
    new_uri += "//" + loc.host;
    new_uri += loc.pathname + "buttonpress";
    return new_uri;
}

let output = document.getElementById("output");
let input = document.getElementById("input");
let ws;

const sendS = () =>{
    if (!ws) {
        return false;
    }
    console.log("SEND: S");
    ws.send("S");
    return false;
};

const sendL = ()=> {
    if (!ws) {
        return false;
    }
    console.log("SEND: L");
    ws.send("L");
    return false;
};

const sendLU = () => {
    if (!ws) {
        return false;
    }
    console.log("SEND: LU");
    ws.send("LU");
    return false;
};

const sendLD = () =>{
    if (!ws) {
        return false;
    }
    console.log("SEND: LD");
    ws.send("LD");
    return false;
};

const sendRU = () =>{
    if (!ws) {
        return false;
    }
    console.log("SEND: RU");
    ws.send("RU");
    return false;
};

const sendRD = ()=>{
    if (!ws) {
        return false;
    }
    console.log("SEND: RD");
    ws.send("RD");
    return false;
};

document.querySelector("body").onload = (evt) => {
    if (ws) {
        return false;
    }
    ws = new WebSocket(getSocketURI());
    ws.onopen = function(evt) {
        console.log("OPEN");
    }
    ws.onclose = function(evt) {
        console.log("CLOSE");
        ws = null;
    }
    ws.onmessage = function(evt) {
        console.log("RESPONSE: " + evt.data);
    }
    ws.onerror = function(evt) {
        console.log("ERROR: " + evt.data);
    }

    document.getElementById("startbutton").onclick = sendS;
    
    document.getElementById("launchbutton").onclick = sendL;
    
    document.getElementById("leftbutton").onmousedown = sendLU;

    document.getElementById("rightbutton").onmousedown = sendRU;
    
    document.getElementById("leftbutton").onmouseup = sendLD;
    
    document.getElementById("rightbutton").onmouseup = sendRD;

    // Sign Up
    document.getElementById("signup-button").addEventListener("click", handleSignUp);
    return false;
};

const handleSignUp = (event) => {
    event.preventDefault();
    
    const name = document.getElementById("name").value;
    const initials = document.getElementById("initials").value;
    
    if (!name.length || !initials.length){
        return;
    }
    document.getElementById("name").value = '';
    document.getElementById("initials").value = '';
    
    signup(name, initials);
}