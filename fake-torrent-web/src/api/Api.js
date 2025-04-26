
import axios from 'axios';

const api = axios.create({
    baseURL: process.env.REACT_APP_API_BASE_URL,
});

export const uploadFile = async (file) => {
    const formData = new FormData();
    formData.append('file', file);

    try {
        const response = await api.post('/upload', formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
            },
        });
        return {
            response: response.data,
            status: 'ok'
        };
    } catch (error) {
        console.error('Error uploading file:', error);
        return {
            response: error,
            status: 'error'
        };
    }
};

export const getFilesList = async () => {
    try {
        const response = await api.get('/getFilesList');
        return {
            response: response.data,
            status: 'ok'
        };
    } catch (error) {
        console.error('Error getting files:', error);
        return {
            response: error,
            status: 'error'
        }
    }
}