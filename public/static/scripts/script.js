(async () => {
    try {
        const MP_PUBLIC_KEY = await GetMPPublicKey();
        if (!MP_PUBLIC_KEY) {
            throw new Error("Failed to retrieve MP_PUBLIC_KEY");
        }
        const mp = new MercadoPago(MP_PUBLIC_KEY);
        const bricksBuilder = mp.bricks();

        const stsPath = window.location.pathname;
        const queryParams = new URLSearchParams(window.location.search);
        const paymentId = queryParams.get('payment_id');
        const externalReference = queryParams.get('external_reference');

        if (stsPath === "/success" && paymentId) {
            fetch('/api/payment/check', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    externalReference: externalReference,
                }),
            })
                .then(response => response.json())
                .then(checkData => {
                    if (checkData.status === 'approved') {
                        document.getElementById('watermarkDiv').style.display = 'block';
                    }
                })
                .catch(error => {
                    console.error('Error verifying payment:', error);
                });
        } else {
            fetch('/api/payment/preference', { method: 'POST' })
                .then(response => response.json())
                .then(data => {
                    bricksBuilder.create("wallet", "wallet_container", {
                        initialization: {
                            preferenceId: data.id,
                        },
                        customization: {
                            texts: {
                                valueProp: 'smart_option',
                            },
                        },
                    });
                });
        }
    } catch (error) {
        console.error("Error initializing MercadoPago:", error.message);
    }
})();

async function generateQRCode() {
console.log('generateQRCode');
    var path = "generate";
    const content = document.getElementById('content').value;
    const qrCodeImage = document.getElementById('qrCodeImage');
    const watermarkCheckbox = document.getElementById('watermarkCheckbox');
    const watermarkFile = document.getElementById('watermarkFile').files[0];
    const formData = new FormData();

    if (!content) {
        alert('Content is required!');
        return;
    }

    formData.append('content', content);


    if (watermarkFile)  {
        formData.append('watermark', watermarkFile);
        path = "generateWM";
    }

    try {
        const response = await fetch("/api/"+path, {
            method: 'POST',
            body: formData
        });

        if (response.ok) {
            const blob = await response.blob();
            qrCodeImage.src = URL.createObjectURL(blob);
            qrCodeImage.style.display = 'block';
        } else {
            const errorData = await response.json();
            alert(errorData.message);
        }
    } catch (error) {
        alert('Failed to generate QR code.');
        console.error('Error:', error);
    }
}

async function GetMPPublicKey() {
    try {
        const response = await fetch("/api/payment/mpInfo");
        if (!response.ok) {
            throw new Error(`HTTP Error: ${response.status}`);
        }
        const data = await response.json();
        if (data.error) {
            throw new Error(data.error);
        }
        return data.MP_PUBLIC_KEY;
    } catch (error) {
        console.error("Error fetching site ID:", error.message);
    }
}