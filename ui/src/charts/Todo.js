import React from "react"
import useFetch from "./useFetch";

const Todo = () => {

    const [response, isLoading] = useFetch('http://localhost:8080/v1/api/new_cases');
    
    return (
        <div>
            <p>Todo List</p>
            {!isLoading && response && response.data.map((record, index) => {
                return <p>{index}, {record.year}</p>
            })}
            <hr />
        </div>
    )
}

export default Todo;
