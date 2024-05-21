document.addEventListener("DOMContentLoaded", function() {
    function checkResetpwNull(password, again) {
        return password !== "" && again !== "";
    }
    
    function checkResetpwPassword(password) {
        return password.length >= 4;
    }
    
    function checkResetpwAgain(password, again) {
        return password == again;
    }
    
    function checkResetpw() {
        var password = document.getElementById("password").value;
        var again = document.getElementById("password-again").value;
    
        if (!checkResetpwNull(password, again)) {
            alert("欄位空白");
            return false;
        }
    
        if (!checkResetpwPassword(password)) {
            alert("密碼長度需大於等於4，請重新輸入");
            return false;
        }
    
        if (!checkResetpwAgain(password, again)) {
            alert("密碼不相同，請重新輸入。");
            return false;
        }
        return true;
    }

    document.querySelector('.resetpw-box').addEventListener('submit', function(event) {
        event.preventDefault();

        if (!checkResetpw()) {
            return;
        }
        this.submit();
    });
});