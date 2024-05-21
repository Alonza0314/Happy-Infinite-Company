let login = document.getElementById('login');
let signup = document.getElementById('signup');
let form_box = document.getElementsByClassName('form-box')[0];
let login_box = document.getElementsByClassName('login-box')[0];
let signup_box = document.getElementsByClassName('signup-box')[0];

function isDesktop() {
    return window.innerWidth > 600;
}

if (isDesktop()) {
    login.addEventListener('click', () => {
        form_box.style.transform = 'translateX(0%)';
        signup_box.classList.add('hidden');
        login_box.classList.remove('hidden');
    });
    
    signup.addEventListener('click', () => {
        form_box.style.transform = 'translateX(87%)';
        login_box.classList.add('hidden');
        signup_box.classList.remove('hidden');
    });
}
else {
    login.addEventListener('click', () => {
        signup_box.classList.add('hidden');
        login_box.classList.remove('hidden');
    });
    
    signup.addEventListener('click', () => {
        login_box.classList.add('hidden');
        signup_box.classList.remove('hidden');
    });
}

document.addEventListener("DOMContentLoaded", function() {
    function checkSignupNull(username, email, password, passwordAgain) {
        return username !== "" && email !== "" && password !== "" && passwordAgain !== ""
    }
    
    function checkSignupUsername(username) {
        var regex = /^(?=.*\d)[a-z\d]+$/;
        return regex.test(username);
    }
    
    function checkSignupEmail(email) {
        var regex = /\S+@\S+\.\S+/;
        return regex.test(email);
    }
    
    function checkSignupPassword(password) {
        return password.length >= 4;
    }
    
    function checkSignupPasswordAgain(password, passwordAgain) {
        return password == passwordAgain;
    }
    
    function checkSignup() {
        var username = document.getElementById("username").value;
        var email = document.getElementById("email").value;
        var password = document.getElementById("password").value;
        var passwordAgain = document.getElementById("password-again").value;
    
        if (!checkSignupNull(username, email, password, passwordAgain)) {
            alert("欄位空白");
            return false;
        }
    
        if (!checkSignupUsername(username)) {
            alert("使用者名稱僅能且必須包含小寫英文及數字，請重新輸入");
            return false;
        }
    
        if (!checkSignupEmail(email)) {
            alert("電子郵件格式錯誤，請重新輸入");
            return false;
        }
    
        if (!checkSignupPassword(password)) {
            alert("密碼長度需大於等於4，請重新輸入");
            return false;
        }
    
        if (!checkSignupPasswordAgain(password, passwordAgain)) {
            alert("密碼不相同，請重新輸入");
            return false;
        }
    
        return true;
    }
    
    function checkLoginNull() {
        var username = document.getElementById("login-username").value;
        var password = document.getElementById("login-password").value;
    
        if (username == "" || password == "") {
            alert("欄位空白");
            return false;
        }
        return true;
    }
    
    document.querySelector('.signup-box').addEventListener('submit', function(event) {
        event.preventDefault();

        if (!checkSignup()) {
            return;
        }
        this.submit();
    });

    document.querySelector('.login-box').addEventListener('submit', function(event) {
        event.preventDefault();

        if (!checkLoginNull()) {
            return;
        }
        this.submit();
    });
});

function getUrlParameter(name) {
    name = name.replace(/[\[]/, '\\[').replace(/[\]]/, '\\]');
    var regex = new RegExp('[\\?&]' + name + '=([^&#]*)');
    var results = regex.exec(location.search);
    return results === null ? '' : decodeURIComponent(results[1].replace(/\+/g, ' '));
}

var signupCondition= getUrlParameter('signup');
if (signupCondition === 'true') {
    alert('註冊成功\n\n請重新登入');
} else {
    if (signupCondition !== '') {
        alert('註冊失敗：' + signupCondition);
    }
}

var loginCondition = getUrlParameter('login');
if (loginCondition !== '') {
    alert('登入失敗：' + loginCondition);
}

var resetpwCondition = getUrlParameter('resetpw');
if (resetpwCondition !== '') {
    alert(resetpwCondition);
}