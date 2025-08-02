document.addEventListener("DOMContentLoaded", () => {
    // Selección de elementos del DOM
    const toogleButton = document.querySelector(".navbar-toogle-btn");
    const mobileMenu = document.querySelector(".navbar-mobile-menu");

    // Si el menu movil está oculto ('none' o vacio), lo muestra cambiando a 'flex'
    // Si el menu movil ya está visible ('flex'), lo oculta cambiando a 'none'

    const toogleMenu = () => {
        mobileMenu.style.display =
            mobileMenu.style.display === "none" || mobileMenu.style.display === ""
                ? "flex"
                : "none";
    };

    const hideMenuResize = () => {
        mobileMenu.style.display = "none"
    }

    // Event Listener
    toogleButton.addEventListener("click", toogleMenu);
    window.addEventListener("resize", hideMenuResize);
    window.addEventListener("load", hideMenuResize);

}); 