import React, { useState } from 'react';

interface VoteProps {
    postid: number;
}

const VoteComponent: React.FC<VoteProps> = (props: VoteProps) => {

    const [value, setValue] = useState(0);

    const voteButtonClicked = (e: any, buttonValue: number) => {
        e.preventDefault();

        if(buttonValue == value) buttonValue = 0;

        if(buttonValue == 1)
        {
            let res = fetch("/api/post/"+props.postid+"/vote/like", {
                method: "post"
            });
        }
        else if(buttonValue == -1)
        {
            let res = fetch("/api/post/"+props.postid+"/vote/dislike", {
                method: "post"
            });
        }
        else if(buttonValue == 0)
        {
            let res = fetch("/api/post/"+props.postid+"/vote", {
                method: "delete"
            });
        }

        setValue(buttonValue);
    };

    return (
        <div className='vote'>
            <button onClick={(e: any) => voteButtonClicked(e, 1)}>^</button>
            <button onClick={(e: any) => voteButtonClicked(e, -1)}>v</button>
        </div>
    );
};

export default VoteComponent;