const  handleErrors = (response) => {
    if (!response.ok) {
        return response.text().then(text => {throw new Error(text)})
    }
    return response;
}

const signup = (name,initials) => {
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
        .then(handleErrors)
        .then(response => response.json())
        .then(result => getList(result))
        .catch(error => console.log('error', error.json()));
}

const parseJwt = (token) => {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
};

const getList = (result) => {
    console.log("result: ", result);
    window.localStorage.setItem('user',result.id);
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    const requestOptions = {
        method: 'GET',
        headers: myHeaders,
        credentials: 'include',
        redirect: 'follow'
    };
    
    fetch("/game", requestOptions)
        .then(handleErrors)
        .then(response => response.json())
        .then(result => updateList(result))
        .catch(error => console.log('error', error));
    
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
        if (evt.data === "NEW TURN"){
            const myHeaders = new Headers();
            myHeaders.append("Content-Type", "application/json");
            const requestOptions = {
                method: 'GET',
                headers: myHeaders,
                credentials: 'include',
                redirect: 'follow'
            };
            
            fetch("/game", requestOptions)
                .then(handleErrors)
                .then(response => response.json())
                .then(result => updateList(result))
                .catch(error => console.log('error', error));
        }
    }
    ws.onerror = function(evt) {
        console.log("ERROR: " + evt.data);
    }

    fetch("/controls", )
    .then(handleErrors)
    .then(response => response.json())
    .then(result => bindHandlers(result))
    .catch(error => console.log('error', error));

    const bindHandlers = (ButtonCommands) => {
        ButtonCommands.map((com) => {
            if (com.dom_id.length){
                if (com.down_command.length){
                    if (!com.up_command.length){
                        document.getElementById(com.dom_id).onclick = () => {wsSend(com.down_command)};
                    }else{
                        document.getElementById(com.dom_id).onmousedown = () => {wsSend(com.down_command)};
                    }
                }
                if (com.up_command.length){
                    document.getElementById(com.dom_id).onmouseup = ()=> {wsSend(com.up_command)};
                }
            }
        });


        document.onkeyup = (e) => {
            e = e || window.event;
            ButtonCommands.forEach((com)=>{
                if(com.keys.indexOf(e.key)>-1){
                    wsSend(com.down_command);
                }
            });
        };

        document.onkeydown = (e) => {
            e = e || window.event;
            ButtonCommands.forEach((com)=>{
                if(com.keys.indexOf(e.key)>-1){
                    wsSend(com.up_command);
                }
            });
        };
    }
}

const refreshList = (userID) => {
    if (!userID || !userID.length){
        return;
    }
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    const requestOptions = {
        method: 'GET',
        headers: myHeaders,
        credentials: 'include',
        redirect: 'follow'
    };
    
    fetch("/game", requestOptions)
        .then(handleErrors)
        .then(response => response.json())
        .then(result => updateList(result))
        .catch(error => console.log('error', error));
}

const updateList = (result) => {
    const queueList = document.getElementById('queue-list');
    
    let arr = [];
    for (const element of result) {
        let listItem = document.createElement('li');
        listItem.textContent = `${element.score || element.name} - ${element.initials}`;
        arr.push(listItem);
        listItem.classList.add(element.class);
    }
    
    queueList.replaceChildren(...arr);
    document.getElementById("signin-button").classList.add("remove");
    document.getElementById("signup-button").classList.add("remove");
    document.getElementById("logout-button").classList.remove("remove");
    enableScroll();
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
        if (evt.data === "NEW TURN"){
            const myHeaders = new Headers();
            myHeaders.append("Content-Type", "application/json");
            const requestOptions = {
                method: 'GET',
                headers: myHeaders,
                credentials: 'include',
                redirect: 'follow'
            };
            
            fetch("/game", requestOptions)
                .then(handleErrors)
                .then(response => response.json())
                .then(result => updateList(result))
                .catch(error => console.log('error', error));
        }
    }
    ws.onerror = function(evt) {
        console.log("ERROR: " + evt.data);
    }
}

const signin = (name,initials) => {
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

    fetch("/signin", requestOptions)
    .then(handleErrors)
    .then(response => response.text())
    .then(result => getList(result))
    .catch(error => console.log('error', error));
}

const myHeaders = new Headers();

const requestOptions = {
  method: 'GET',
  headers: myHeaders,
  redirect: 'follow'
};

const getSocketURI = () => {
    let loc = window.location, new_uri;
    if (loc.protocol === "https:") {
        new_uri = "wss:";
    } else {
        new_uri = "ws:";
    }
    new_uri += "//" + loc.host;
    new_uri += loc.pathname + "buttonpress/" + window.localStorage.getItem("user");
    return new_uri;
}

let output = document.getElementById("output");
let input = document.getElementById("input");
let ws;

const wsSend = (val) => {
    if (!ws || !val.length){
        return false;
    }
    console.log(`SEND: ${val}`);
    ws.send(val);
    return false;
}

document.querySelector("body").onload = (evt) => {
    refreshList(window.localStorage.getItem("user"));
    // Sign Up
    document.getElementById("signup-button").addEventListener("click", handleSignUp);
    document.getElementById("signin-button").addEventListener("click", handleSignIn);
    document.getElementById("logout-button").addEventListener("click", handleLogout);
    fetch("/game", requestOptions)
    .then(handleErrors)
    .then(response => response.json())
    .then(result => updateList(result))
    .catch(error => console.log('error', error));
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
    document.getElementById("signin-button").classList.add("remove");
    document.getElementById("signup-button").classList.add("remove");
    document.getElementById("logout-button").classList.remove("remove");
    
    signup(name, initials);
}

const handleSignIn = (event) => {
    event.preventDefault();
    
    const name = document.getElementById("name").value;
    const initials = document.getElementById("initials").value;
    
    if (!name.length || !initials.length){
        return;
    }
    document.getElementById("name").value = '';
    document.getElementById("initials").value = '';
    document.getElementById("signin-button").classList.add("remove");
    document.getElementById("signup-button").classList.add("remove");
    document.getElementById("logout-button").classList.remove("remove");
    signin(name, initials);
}

const handleLogout = (event) => {
    event.preventDefault();
    
    document.getElementById("name").value = '';
    document.getElementById("initials").value = '';

    // window.localStorage.Item("user")

    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    
    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        redirect: 'follow'
    };
    
    fetch("/logout", requestOptions)
        .then(response => response.json())
        .then(result => logout(result))
        .catch(error => console.log('error', error));
    document.getElementById("signin-button").classList.remove("remove");
    document.getElementById("signup-button").classList.remove("remove");
    document.getElementById("logout-button").classList.add("remove");
}

const logout = (result) =>{
    window.localStorage.removeItem("user");
}

const showLowerSections = () => {
    const queue = document.getElementsByClassName('queue');
    const cabinet = document.getElementsByClassName('cabinet');
    queue[0].classList.remove('hidden-section');
    cabinet[0].classList.remove('hidden-section');
  };
  
const disableScroll = () => {
  window.onscroll = () => {
    window.scroll(0, 0);
  };
};
  
disableScroll();
  
const enableScroll = () => {
  window.onscroll = function () {};
  showLowerSections();
  
  window.scrollBy({
    top: window.innerHeight,
    behavior: 'smooth',
  });
};
