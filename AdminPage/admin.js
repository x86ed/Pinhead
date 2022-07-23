import { deleteUser } from "./api.js";

export const confirmDeleteUser = (username, userId) => {
    document.getElementById("deleteUserText").textContent = "Are you sure you want to delete " + username + ":"; 
    document.getElementById("confirmDeleteUserBtn").value = userId + ":user";
    document.getElementById('modal-delete').open = true;
}

export const confirmDeleteAdminUser = (email, userId) => {
    document.getElementById("deleteAdminUserText").textContent = "Are you sure you want to delete " + email + ":"; 
    document.getElementById("confirmDeleteAdminUserBtn").value = userId + ":admin";
    document.getElementById('modal-delete-admin').open = true;
}

const deleteUserBtn = () => {
    const deleteUserId = document.getElementById("confirmDeleteUserBtn").value;
    console.log("Delete userId: ", deleteUserId);
    deleteUser(deleteUserId);
    document.getElementById("confirmDeleteUserBtn").value = undefined;
    document.getElementById('modal-delete').open = false;
}

const deleteAdminUserBtn = () => {
    const deleteUserId = document.getElementById("confirmDeleteAdminUserBtn").value;
    console.log("Delete userId: ", deleteUserId);
    deleteUser(deleteUserId);
    document.getElementById("confirmDeleteAdminUserBtn").value = undefined;
    document.getElementById('modal-delete-admin').open = false;
}

document.querySelector("body").onload = (evt) => {
    document.getElementById("confirmDeleteUserBtn").addEventListener('click', deleteUserBtn);
    document.getElementById("confirmDeleteAdminUserBtn").addEventListener('click', deleteAdminUserBtn);
    return false;
}