import { useEffect, useState } from "react"


const useFetch = (url) => {

    const [isLoading, setIsLoading] = useState(true);
    const [hasError, setHasError] = useState(false);
    const [response, setResponse] = useState([]);

    useEffect(() => {
        setIsLoading(true);
        
        fetch(url)
        .then((res) => res.json())
        .then((response) => {
            setIsLoading(false);
            setResponse(response)
        })
        .catch((err) => setHasError(true));
// eslint-disable-next-line
    }, []);

    return [response, isLoading, hasError];

}

export default useFetch;