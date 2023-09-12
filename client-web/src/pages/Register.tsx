import React from 'react';
import { useNavigate } from 'react-router';
import * as bcrypt from 'bcryptjs';

interface RegisterProps {
    refetchUser: () => void;
}

const Register: React.FC<RegisterProps> = (props: RegisterProps) => {
    const navigate = useNavigate();

    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        let formData = new FormData(e.currentTarget);

        var password_plain = formData.get("password") as string;

        var salt = bcrypt.genSaltSync(10);
        var password_hash = bcrypt.hashSync(formData.get("password") as string);
        formData.set("password", password_hash);

        let res = fetch("/api/user", {
            body: formData,
            method: "post"
        }).finally(() => {
            formData.set("password", password_plain);
            let res = fetch("/api/session", {
                body: formData,
                method: "post"
            })
            .then((r) => r.text())
            .then((text) => {
                // document.cookie = "session="+text;
                navigate('/');
                props.refetchUser();
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
                <input type="password" id="password" name="password" />
                </label>
                <br />
                <input type="submit" value="Register!" />
            </form>
        </div>
    );
}

export default Register;