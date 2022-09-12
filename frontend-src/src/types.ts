enum DownloadStatus {
    Available,
	Download,
	Decryption,
    Done,
}

enum UploadStatus {
    Encryption,
    Upload,
    Done,
}

interface Upload {
    id: string;
    name: string;
    speed: number;
    progress: number;
    status: UploadStatus;
}

interface Download {
    id: string;
    name: string;
    speed: number;
    progress: number;
    status: DownloadStatus;
}

export {
    DownloadStatus,
    UploadStatus,
    Download,
    Upload,
}
