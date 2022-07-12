const baseUrl = "http://localhost:8080/";

const defaultHeaders = (includeAuth) => {
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    if (includeAuth) {
        myHeaders.append("Authorization", "Bearer " + sessionStorage.getItem('jwtToken'));        
    }

    return myHeaders;
}

export const signup = (name, initials) => {
    const raw = JSON.stringify({
      "name": name,
      "password": name+initials,
      "role": "user",
      "initials": initials
    });

    const requestOptions = {
    method: 'POST',
    headers: defaultHeaders(false),
    body: raw,
    redirect: 'follow'
    };

    fetch(baseUrl + "signup", requestOptions)
    .then(response => response.text())
    .then(result => console.log(result))
    .catch(error => console.log('error', error));
}

export async function signin(name, initials) {
    await adminSignin(name, name + initials);
}

export async function adminSignin(name, password) {
    const raw = JSON.stringify({
        "name": name,
        "password": password,
    });

    const requestOptions = {
        method: 'POST',
        headers: defaultHeaders(false),
        body: raw,
        redirect: 'follow'
    };

    try {
        var response = await fetch(baseUrl + "signin", requestOptions)
        var result = await response.text();
    
        console.log(result);
        var authUser = JSON.parse(result);
        console.log("authUser: ", authUser);
        sessionStorage.setItem("userName", authUser.name);
        sessionStorage.setItem("userRole", authUser.role);
        sessionStorage.setItem("jwtToken", authUser.token);
        return authUser;
    }
    catch (error) {
        console.log('error', error);
    };
}
export async function getUsers() {
    const requestOptions = {
    method: 'GET',
    headers: defaultHeaders(true),
    redirect: 'follow'
    };

    try {
        var response = await fetch(baseUrl + "users", requestOptions)
        var result = await response.text();
    
        console.log(result);
        var authUser = JSON.parse(result);
        console.log("authUser: ", authUser);
        return authUser;
    }
    catch (error) {
        console.log('error', error);
    };
}

export const deleteUser = (userId) => {

    const requestOptions = {
        method: 'DELETE',
        headers: defaultHeaders(true),
        redirect: 'follow'
    };
  
    fetch(baseUrl + "admin/" + userId, requestOptions)
    .then(response => response.text())
    .then(result => console.log(result))
    .catch(error => console.log('error', error));
}

export const getSocketURI = () => {
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