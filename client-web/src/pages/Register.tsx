import React from 'react';

export default function Register() {
    return (
        <div>
            <h2>Register</h2>

            <form method="post" action="/api/user/">
                <label>
                Username:
                <input type="text" id="username" name="username" />
                </label>

                <label>
                Password:
                <input type="text" id="password" name="password" />
                </label>

                <input type="submit" value="Register!" />
            </form>
        </div>
    );
} 