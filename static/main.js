import * as api from "/api.js"

/* 
* Admin Modal Controls 
*/
const openAdminSignin = () => {
    document.getElementById("adminModal").style.display = "block";
}

const closeAdminSignin = () => {
    document.getElementById("adminModal").style.display = "none";
    document.getElementById("modalError1").style.visibility = "hidden";
    document.getElementById("modalError1").style.visibility = "hidden";
}

async function submitAdminCred() {
    var username = document.getElementById("adminUsername").value;
    var password = document.getElementById("adminPassword").value;

    if (!!username && !!password) {
        document.getElementById("modalError1").style.visibility = "hidden";
        var authUser = await api.adminSignin(username, password);
        if (authUser && authUser.role === "admin") {
            document.getElementById("modalError2").style.visibility = "hidden";
            window.location.href = 'AdminPage/admin.html';
        }
        else {
            document.getElementById("modalError2").style.visibility = "visible";
        }
    }
    else {
        document.getElementById("modalError1").style.visibility = "visible";
    }
}

//unused???
var output = document.getElementById("output");
var input = document.getElementById("input");

/*
* Websockets
*/
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

document.querySelector("body").onload = function(evt) {
    if (ws) {
        return false;
    }
    ws = new WebSocket(api.getSocketURI());
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

    document.getElementById("openAdminSignin").onclick = openAdminSignin;

    document.getElementById("aModalCancel").onclick = closeAdminSignin;

    document.getElementById("aModalSubmit").onclick = async () => { await submitAdminCred() };

    var modal = document.getElementById("adminModal");
    // When the user clicks anywhere outside of the modal, close it
    window.onclick = function(event) {
    if (event.target == modal) {
        closeAdminSignin();
    }
  }
    return false;
};

/*
these look unused
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
*/