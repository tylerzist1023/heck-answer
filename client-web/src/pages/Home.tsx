import React from 'react';
import {useState, useEffect} from 'react';
import ThreadComponent from '../components/ThreadComponent';

export default function Home() {

    const [threads, setThreads] = useState([]);

    useEffect(() => {
        let res = fetch("/api/threads/", {
            method: "get"
        })
        .then((r) => r.json())
        .then((json) => {
            setThreads(json);
        });
    }, [])

    return (
        <div>
        {
            threads.map((x) => { return (<ThreadComponent 
                id={x["id"]}
                title={x["title"]}
                author={x["userid"]}
                url={x["url"]}
            />);})
        }
        </div>
    );
}