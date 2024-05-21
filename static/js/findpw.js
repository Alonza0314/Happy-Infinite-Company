document.addEventListener("DOMContentLoaded", function() {
    const authImg = document.getElementById("captcha");
    const refresh = document.getElementById("refresh-captcha");
    let correctAuthCode = "";

    function refreshAuthImg() {
        const authUrl = authImg.dataset.authUrl;
        fetch(authUrl)
            .then(response => response.json())
            .then(data => {
                authImg.src = `${data.image}`;
                correctAuthCode = data.code;
            })
            .catch(error => console.error('Error fetching AuthCode:', error));
    }

    refreshAuthImg();

    refresh.addEventListener("click", function() {
        refreshAuthImg();
    });

    function checkFindpwNull(username, email, authcode) {
        return username !== "" && email !== "" && authcode !== "";
    }
    
    function checkFindpwUsername(username) {
        var regex = /^(?=.*\d)[a-z\d]+$/;
        return regex.test(username);
    }
    
    function checkFindpwEmail(email) {
        var regex = /\S+@\S+\.\S+/;
        return regex.test(email);
    }
    
    function checkFindpw() {
        var username = document.getElementById("username").value;
        var email = document.getElementById("email").value;
        var authcode = document.getElementById("captcha-code").value
    
        if (!checkFindpwNull(username, email, authcode)) {
            alert("欄位空白");
            return false;
        }
    
        if (!checkFindpwUsername(username)) {
            alert("使用者名稱僅能且必須包含小寫英文及數字，請重新輸入");
            return false;
        }
    
        if (!checkFindpwEmail(email)) {
            alert("電子郵件格式錯誤，請重新輸入");
            return false;
        }

        return true;
    }

    document.querySelector('.findpw-box').addEventListener('submit', function(event) {
        event.preventDefault();

        if (!checkFindpw()) {
            return;
        }

        const authCodeInput = document.getElementById("captcha-code").value;

        if (correctAuthCode === authCodeInput) {
            this.submit();
        } else {
            alert("驗證碼錯誤，請重新輸入");
            authCodeInput.value = "";
            refreshAuthImg();
            return
        }
    });
});

function getUrlParameter(name) {
    name = name.replace(/[\[]/, '\\[').replace(/[\]]/, '\\]');
    var regex = new RegExp('[\\?&]' + name + '=([^&#]*)');
    var results = regex.exec(location.search);
    return results === null ? '' : decodeURIComponent(results[1].replace(/\+/g, ' '));
}

var findpwCondition= getUrlParameter('findpw');
if (findpwCondition !== '') {
    alert('找回密碼失敗：' + findpwCondition);
}
