import { deleteUser } from "./api.js";

export const confirmDeleteUser = (username, userId) => {
    document.getElementById("deleteUserText").textContent = "Are you sure you want to delete " + username + ":"; 
    document.getElementById("confirmDeleteUserBtn").value = userId;
    document.getElementById('modal-delete').open = true;
}

const deleteUserBtn = () => {
    const deleteUserId = document.getElementById("confirmDeleteUserBtn").value;
    console.log("Delete userId: ", deleteUserId);
    deleteUser(deleteUserId);
    document.getElementById("confirmDeleteUserBtn").value = undefined;
    document.getElementById("deleteModal").style.display = "none";
}

document.querySelector("body").onload = (evt) => {
    document.getElementById("confirmDeleteUserBtn").addEventListener('click', deleteUserBtn);
    return false;
}