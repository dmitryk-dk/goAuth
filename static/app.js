async function post (url, body) {
    const headers = {
        'X-Requested-With': 'XMLHttpRequest',
        'Content-Type':     'application/json',
        'Accept':           'application/json',
    };
    const response = await fetch(
        url,
        {
            method: 'post',
            credentials: 'same-origin',
            redirect: 'manual',
            headers: headers,
            body: JSON.stringify(body)
        }
    );
    return await response.json();
};


const init = () => {
    const submitButton = document.querySelector('.js-submit-button');
    const inputEmail = document.querySelector('.js-input-email');
    const inputPassword = document.querySelector('.js-input-password');
    const url = "http://localhost:3000/login";
    submitButton.addEventListener('click', () => {
        const username = inputEmail.value;
        const password = inputPassword.value;
        const body = { username, password };
        post(url, body)
            .then(data => console.log(data))
            .catch(reason => console.log(reason.message))
    });
};

init();
