import React from 'react';

export default function Login() {
    return (
        <div>
            <h2>Login</h2>

            <form method="POST" action="/api/user">
                <label>
                Username:
                <input type="text" id="username" name="username" />
                </label>

                <label>
                Password:
                <input type="text" id="password" name="password" />
                </label>

                <input type="submit" value="Login!" />
            </form>
        </div>
    );
}