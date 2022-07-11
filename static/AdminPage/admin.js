import * as api from "../api.js";

export const confirmDeleteUser = (username, userId) => {
    console.log("name: ", username);
    console.log("id: ", userId);

    document.getElementById("deleteUserText").textContent = "Are you sure you want to delete " + username + ":"; 
    document.getElementById("confirmDeleteUserBtn").value = userId;
    document.getElementById("deleteModal").style.display = "block";
}

const cancelDelete = () => {
    document.getElementById("confirmDeleteUserBtn").value = undefined;
    document.getElementById("deleteModal").style.display = "none";
}

const deleteUser = () => {
    const deleteUserId = document.getElementById("confirmDeleteUserBtn").value;
    console.log("Delete userId: ", deleteUserId);
    api.deleteUser(deleteUserId);
}

document.getElementById("confirmDeleteUserBtn").addEventListener('click', deleteUser);
document.getElementById("cancelDeleteBtn").addEventListener('click', cancelDelete);