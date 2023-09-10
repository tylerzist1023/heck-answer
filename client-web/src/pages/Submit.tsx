import React from 'react';
import { useNavigate } from 'react-router';

const Submit = () => {
    const navigate = useNavigate();

    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        let formData = new FormData(e.currentTarget);

        let res = fetch("/api/post/", {
            body: formData,
            method: "post"
        }).finally(() => {
            navigate('/');
        });
    };

    return (
        <div>
            <h2>Submit</h2>

            <form method="post" onSubmit={onSubmit}>
                <label>
                Url:
                <input type="text" id="url" name="url" />
                </label>

                <br />

                <label>
                Title:
                <input type="text" id="title" name="title" />
                </label>

                <br />

                <label>
                Body:
                <textarea id="body" name="body" rows={10} cols={100}></textarea>
                </label>

                <br />

                <input type="submit" value="Submit!" />
            </form>
        </div>
    );
}

export default Submit;