import React from 'react';
import {useState, useEffect} from 'react';

const User = () => {
    const queryParams = new URLSearchParams(window.location.search);

    const [userid, setUserId] = useState(parseInt(queryParams.get("id")!));
    const [user, setUser] = useState();

    useEffect(() => {
        let res = fetch("/api/user/"+userid, {
            method: "get"
        })
            .then((r) => r.json())
            .then((json) => {
                setUser(json);
            });
    }, [userid]);

    if(user == undefined) {
        return <div></div>;
    } else {
        return (
            <div>
                <p>id={userid}</p>
                <p>name={user["username"]}</p>
            </div>
        );
    }
}

export default User;