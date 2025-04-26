import React, { useState, useEffect } from 'react';
import FileUpload from './FileUpload';
import { getFilesList } from '../api/Api';
import FileList from './FileList';
import './styles/Main.css'

const Main = () => {
    const [files, setFiles] = useState([]);
    const [filteredFiles, setFilteredFiles] = useState([]);
    const [isLoading, setIsLoading] = useState(true);
    const [showUpload, setShowUpload] = useState(false);
    const [searchQuery, setSearchQuery] = useState('');

    useEffect(() => {
        fetchFiles();
    }, []);

    useEffect(() => {
        const filtered = files.filter(file =>
            file.name.toLowerCase().includes(searchQuery.toLowerCase())
        );
        setFilteredFiles(filtered);
    }, [searchQuery, files]);

    const fetchFiles = async () => {
        setIsLoading(true);
        try {
            const filesList = (await getFilesList()).response;
        } catch (error) {
            console.error('Error fetching files:', error);
            setFiles([]);
        } finally {
            setIsLoading(false);
        }
    };

    const handleFileAction = (fileId) => {
        console.log('Файл:', fileId);
        // TODO
    };

    const handleUploadSuccess = () => {
        setShowUpload(false);
        fetchFiles();
    };

    const handleSearchChange = (e) => {
        setSearchQuery(e.target.value);
    };

    return (
        <div className="main-container">
            {showUpload ? (
                <FileUpload
                    onSuccess={handleUploadSuccess}
                    onCancel={() => setShowUpload(false)}
                />
            ) : (
                <div className='m-win'>
                    <div className="header">
                        <h1>(не) Fake Torrent</h1>
                    </div>
                    <div className="search-container">
                        <input
                            type="text"
                            placeholder="🔍 Поиск файлов"
                            value={searchQuery}
                            onChange={handleSearchChange}
                            className="search-input"
                        />
                    </div>

                    <FileList
                        files={filteredFiles}
                        isLoading={isLoading}
                        onFileClick={handleFileAction}
                    />
                    <button
                        onClick={() => setShowUpload(true)}
                        className="upload-button btn"
                    >
                        💾 Загрузить файл
                    </button>
                </div>
            )}
        </div>
    );
};

export default Main;