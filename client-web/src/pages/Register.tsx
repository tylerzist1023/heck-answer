import React from 'react';
import { useNavigate } from 'react-router';

const Register = () => {
    const navigate = useNavigate();

    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        let formData = new FormData(e.currentTarget);

        let res = fetch("/api/user/", {
            body: formData,
            method: "post"
        }).finally(() => {
            let res = fetch("/api/session/", {
                body: formData,
                method: "post"
            })
            .then((r) => r.text())
            .then((text) => {
                document.cookie = "session="+text;
                navigate('/');
                window.location.reload();
            });
        });
    };

    return (
        <div>
            <h2>Register</h2>

            <form method="post" onSubmit={onSubmit}>
                <label>
                Username:
                <input type="text" id="username" name="username" />
                </label>
                <br />
                <label>
                Password:
                <input type="text" id="password" name="password" />
                </label>
                <br />
                <input type="submit" value="Register!" />
            </form>
        </div>
    );
}

export default Register;