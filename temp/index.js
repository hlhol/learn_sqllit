// index.js

function switchForm() {
    const formTitle = document.getElementById('form-title');
    const authForm = document.getElementById('auth-form');
    const submitButton = authForm.querySelector('button[type="submit"]');
    const switchButton = document.querySelector('.switch-button');
    
    if (formTitle.textContent === 'Login') {
        formTitle.textContent = 'Sign Up';
        authForm.action = '/signup';
        submitButton.textContent = 'Sign Up';
        switchButton.textContent = 'Switch to Login';
    } else {
        formTitle.textContent = 'Login';
        authForm.action = '/login';
        submitButton.textContent = 'Log In';
        switchButton.textContent = 'Switch to Sign Up';
    }
}
