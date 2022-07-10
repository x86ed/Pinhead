let authUser;

const signup = (name, initials) => {
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
    .then(result => console.log(result))
    .catch(error => console.log('error', error));
}

const signin = async (name,password)=>{
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    console.log("signing in user: ", name);

    const raw = JSON.stringify({
    "name": name,
    "password": password //normal case password = name + initials
    });

    const requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: raw,
    redirect: 'follow'
    };

    fetch("http://localhost:8080/signin", requestOptions)
        .then(response => response.text())
        .then(result => {
            console.log(result);
            authUser = JSON.parse(result);
            console.log("authUser: ", authUser);
            sessionStorage.setItem("userName", authUser.name);
            sessionStorage.setItem("userRole", authUser.role);
            sessionStorage.setItem("jwtToken", authUser.TokenString);
        })
        .catch(error => {
            console.log('error', error);
            authUser = null;
        });
}

var myHeaders = new Headers();
myHeaders.append("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTU2Njg5MjAsIm5hbWUiOiJkZXJwIiwicm9sZSI6InVzZXIifQ.w09Rwoa7X0Fu5aMXrvQ5KMwA5VeSpMhSJ1j24snTdJU");

var requestOptions = {
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

var output = document.getElementById("output");
var input = document.getElementById("input");
let ws;

var print = function(message) {
console.log(message);
};

const sendS = () =>{
    if (!ws) {
        return false;
    }
    print("SEND: S");
    ws.send("S");
    return false;
};

const sendL = ()=> {
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

const sendLD = () =>{
    if (!ws) {
        return false;
    }
    print("SEND: LD");
    ws.send("LD");
    return false;
};

const sendRU = () =>{
    if (!ws) {
        return false;
    }
    print("SEND: RU");
    ws.send("RU");
    return false;
};

const sendRD = ()=>{
    if (!ws) {
        return false;
    }
    print("SEND: RD");
    ws.send("RD");
    return false;
};

var adminModal = document.getElementById("myModal");

const openAdminSignin = () => {
    adminModal.style.display = "block";
}

const closeAdminSignin = () => {
    adminModal.style.display = "none";
    document.getElementById("modalError1").style.visibility = "hidden";
    document.getElementById("modalError1").style.visibility = "hidden";
}

const submitAdminCred = async () => {
    var username = document.getElementById("adminUsername").value;
    var password = document.getElementById("adminPassword").value;

    if (!!username && !!password) {
        document.getElementById("modalError1").style.visibility = "hidden";
        console.log("good")
        await signin(username, password);
        if (authUser && authUser.role === "admin") {
            console.log("yay adim");
            document.getElementById("modalError2").style.visibility = "hidden";
            window.location.href = 'AdminPage/admin.html';
        }
        else {
            console.log("NOOOO");
            document.getElementById("modalError2").style.visibility = "visible";
        }
    }
    else {
        document.getElementById("modalError1").style.visibility = "visible";
    }
}

document.querySelector("body").onload = function(evt) {
    if (ws) {
        return false;
    }
    ws = new WebSocket(getSocketURI());
    ws.onopen = function(evt) {
        print("OPEN");
    }
    ws.onclose = function(evt) {
        print("CLOSE");
        ws = null;
    }
    ws.onmessage = function(evt) {
        print("RESPONSE: " + evt.data);
    }
    ws.onerror = function(evt) {
        print("ERROR: " + evt.data);
    }

    document.getElementById("startbutton").onclick = sendS;
    
    document.getElementById("launchbutton").onclick = sendL;
    
    document.getElementById("leftbutton").onmousedown = sendLU;

    document.getElementById("rightbutton").onmousedown = sendRU;
    
    document.getElementById("leftbutton").onmouseup = sendLD;
    
    document.getElementById("rightbutton").onmouseup = sendRD;

    //document.getElementById("adminBtn").onclick = signin;

    
    document.getElementById("adminPage").onclick = openAdminSignin;

    document.getElementById("aModalCancel").onclick = closeAdminSignin;

    document.getElementById("aModalSubmit").onclick = async () => { await submitAdminCred() };

    var modal = document.getElementById("myModal");
    // When the user clicks anywhere outside of the modal, close it
    window.onclick = function(event) {
    if (event.target == modal) {
        closeAdminSignin();
    }
  }


    return false;
};

document.getElementById("send").onclick = function() {
    if (!ws) {
        return false;
    }
    print("SEND: " + input.value);
    ws.send(input.value);
    return false;
};

document.getElementById("close").onclick = function() {
    if (!ws) {
        return false;
    }
    ws.close();
    return false;
};

function parseJwt (token) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
};