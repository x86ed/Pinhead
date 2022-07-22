const baseUrl = "http://localhost:54321/";

const defaultHeaders = (includeAuth) => {
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    if (includeAuth) {
        myHeaders.append("Authorization", "Bearer " + sessionStorage.getItem('jwtToken'));        
    }

    return myHeaders;
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
        var authUser = JSON.parse(result);
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
        var authUser = JSON.parse(result);
        return authUser;
    }
    catch (error) {
        console.log('error', error);
    };
}

export async function getAdmins() {
    const requestOptions = {
    method: 'GET',
    headers: defaultHeaders(true),
    redirect: 'follow'
    };

    try {
        var response = await fetch(baseUrl + "admins", requestOptions)
        var result = await response.text();
        var authUser = JSON.parse(result);
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