//Admin JS support
const getdata = (username, userId) => {
    console.log("name: ", username);
    console.log("id: ", userId);

    let m = document.getElementById("deleteModal");
    document.getElementById("deleteUserText").textContent = "Are you sure you want to delete " + username + " user:"; 
    let mb = document.getElementById("confirmDeleteUserBtn");

    mb.value = userId;
    m.style.display = "block";
}

const cancelDelete = () => {
    document.getElementById("confirmDeleteUserBtn").value = undefined;
    document.getElementById("deleteModal").style.display = "none";
}
const deleteUser = (userId) => {
    console.log("Deltee userId: ", userId);
}

// Admin API
const getUsers = () => {
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTU2Njg5MjAsIm5hbWUiOiJkZXJwIiwicm9sZSI6InVzZXIifQ.w09Rwoa7X0Fu5aMXrvQ5KMwA5VeSpMhSJ1j24snTdJU");

    const requestOptions = {
    method: 'GET',
    headers: myHeaders,
    redirect: 'follow'
    };

    fetch("http://localhost:8080/users", requestOptions)
        .then(response => response.text())
        .then(result => {
            console.log(result);
            authUser = JSON.parse(result);
            console.log("authUser: ", authUser);
        })
        .catch(error => {
            console.log('error', error);
            authUser = null;
        });
}