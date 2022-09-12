export function formatSize(bytes: number): string {
    if (bytes === 0) {
        return "0 Byte"; // TODO: use translation
    }

    const sizes = ["Ko", "Mo", "Go", "To"];
    const i = Math.floor(Math.log(bytes) / Math.log(1024));

    return parseFloat((bytes / Math.pow(1024, i)).toLocaleString(undefined, { maximumFractionDigits: 2 })) + ' ' + sizes[i];
}
