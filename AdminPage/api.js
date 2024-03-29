const baseUrl = `http://${window.location.host}/`;

const  handleErrors = (response) => {
    if (!response.ok) {
        throw Error(response);
    }
    return response;
}

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
        "email": name,
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
        sessionStorage.setItem("userEmail", authUser.email);
        sessionStorage.setItem("jwtToken", authUser.token);
        console.log("got authUser: ", authUser);
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

export async function getGame() {
    const requestOptions = {
    method: 'GET',
    headers: defaultHeaders(true),
    redirect: 'follow'
    };

    try {
        var response = await fetch(baseUrl + "game", requestOptions)
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

export const deleteUser = (userInfo) => {
    const requestOptions = {
        method: 'DELETE',
        headers: defaultHeaders(true),
        redirect: 'follow'
    };
  
    const user = userInfo.split(":");

    fetch(baseUrl + "admin" + "?userId=" + user[0] + "&role=" + user[1], requestOptions)
    .then(response => response.text())
    .then(result => console.log(result))
    .catch(error => console.log('error', error));
}

export async function newGame() {
    const requestOptions = {
        method: 'POST',
        headers: defaultHeaders(true),
        redirect: 'follow'
    };

    try {
        var response = await fetch(baseUrl + "new_game", requestOptions)
        var result = await response.text();
        var resultJson = JSON.parse(result);

        console.log("newGame resultJson: ", resultJson);
        return resultJson;
    }
    catch (error) {
        console.log('error', error);
    };
}

function setUsersResponse(result){
        console.log("nextTurn result: ", result);
        result.forEach(element => {
                window.localStorage.setItem('currentUser', element.user);
        });
}

export async function nextTurn() {
    const requestOptions = {
        method: 'POST',
        headers: defaultHeaders(true),
        redirect: 'follow'
    };
    fetch(baseUrl + "next_turn", requestOptions)
    .then(handleErrors)
    .then(response => response.json())
    .then(result => setUsersResponse(result))
    .catch(error => console.log('error', error.json()));
}

function hScore(resultJson){
    console.log("highScore resultJson: ", resultJson);
    resultJson.forEach(element => {
        if (element.active && !element.complete){
            window.localStorage.setItem('currentUser', element.user);
        }
    });
}

export async function highScore() {
    const requestOptions = {
        method: 'POST',
        headers: defaultHeaders(true),
        redirect: 'follow'
    };
    fetch(baseUrl + "high_score", requestOptions)
    .then(handleErrors)
    .then(response => response.json())
    .then(result => hScore(result))
    .catch(error => console.log('error', error.json()));
}

export async function updateScore(userId, score) {
    const raw = JSON.stringify({
        "id": userId,
        "score": score,
    });

    const requestOptions = {
        method: 'POST',
        headers: defaultHeaders(true),
        body: raw,
        redirect: 'follow'
    };
    fetch(baseUrl + "update_score", requestOptions)
    .then(handleErrors)
    .then(response => response.json())
    .then(result => hScore(result))
    .catch(error => console.log('error', error.json()));
}
