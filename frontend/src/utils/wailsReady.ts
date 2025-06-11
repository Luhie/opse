export async function waitWailsReady(): Promise<void> {
    return new Promise((resolve, reject) => {
        const interval = setInterval(() => {
            if (window.go && window.go.main) {
                clearInterval(interval);
                resolve();
            }
        }, 100);

        // 타임아웃 방지 (10초)
        setTimeout(() => {
            clearInterval(interval);
            reject(new Error("Wails API not initialized"));
        }, 10000);
    });
}
