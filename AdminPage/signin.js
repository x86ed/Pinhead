import { adminSignin } from "./api.js";

async function signinUser() {
    var username = document.getElementById("adminUsername").value;
    var password = document.getElementById("adminPassword").value;

    if (!!username && !!password) {
        document.getElementById("sError1").style.visibility = "hidden";
        var authUser = await adminSignin(username, password);
        if (authUser && authUser.role === "admin") {
            document.getElementById("sError2").style.visibility = "hidden";
            window.location.href = 'index.html';
        }
        else {
            document.getElementById("sError2").style.visibility = "visible";
        }
    }
    else {
        document.getElementById("sError1").style.visibility = "visible";
    }
}

document.querySelector("body").onload = (evt) => {
    document.getElementById("aModalSubmit").addEventListener('click', async () => await signinUser());
    return false;
}