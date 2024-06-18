// index.js

function switchForm() {
    const formTitle = document.getElementById('form-title');
    const authForm = document.getElementById('auth-form');
    const signupForm = document.getElementById('signup-form');
    const switchButton = document.querySelector('.switch-button');
    
    if (formTitle.textContent === 'Login') {
        formTitle.textContent = 'Sign Up';
        authForm.style.display = 'none';
        signupForm.style.display = 'block';
        switchButton.textContent = 'Switch to Login';
    } else {
        formTitle.textContent = 'Login';
        authForm.style.display = 'block';
        signupForm.style.display = 'none';
        switchButton.textContent = 'Switch to Sign Up';
    }
}
