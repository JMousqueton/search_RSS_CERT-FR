---
Title: Veille ciblée avis CERT-FR
Group: Veille
Author: m8Ae4GcE
Creation-Date: 04/02/2022
Status: Version 1
---

## Références 
- https://www.cert.ssi.gouv.fr/avis/feed/
- https://pkg.go.dev/net/http
- https://pkg.go.dev/encoding/xml

# Veille des flux RSS des avis de sécurité du CERT-FR via une liste de mots clés
Permettre de faire de la veille des avis du CERT-FR sur des technologies spécifiques

## Pre-Requis
- GNU/Linux ou Windows x86 minimum
- Port HTTPS ouvert

## Usage

1. Mettre les mots clés à rechercher dans le fichier *search.conf*, exemple :
```
Debian
Ubuntu
VirtualBox
Windows
Android
Nextcloud
Firefox
```

### 2. Linux :

```
./request_rss_cert-fr
```

### 2. Windows :

Double clic sur l'executable ou ouverture du script dans un invite de commande.

## Resultats
CERTFR-2022-AVI-109;Multiples vulnérabilités dans le noyau Linux d’Ubuntu;04 février 2022;https://www.cert.gouv.fr/avis/CERTFR-2022-AVI-109/

CERTFR-2022-AVI-087;Vulnérabilité dans Nextcloud pour Android;27 janvier 2022;https://www.cert.ssi.gouv.fr//CERTFR-2022-AVI-087/

CERTFR-2022-AVI-085;Vulnérabilité dans pkexec de PolicyKit sur Debian;27 janvier 2022;https://www.cert.ssi.fr/avis/CERTFR-2022-AVI-085/

CERTFR-2022-AVI-083;Vulnérabilité dans pkexec de PolicyKit sur Ubuntu;27 janvier 2022;https://www.cert.ssi.fr/avis/CERTFR-2022-AVI-083/

CERTFR-2022-AVI-074;Vulnérabilité dans Microsoft Windows;24 janvier 2022;https://www.cert.ssi.gouv.fr/avis/FR-2022-AVI-074/

Des lignes en double apparaissent en cas de multiples matches

## Intégrité (SHA256)
cce6bad358113c93cceb69f2909a9b3fb93c1e1e5dc878e0b18da3f792b4a1c2  request_rss_cert-fr
99d8379352cbd9e7c6dbfe29f3a981cf97cd02117f9c3f1ac854f20c8032ea37  request_rss_cert-fr.exe
53994e5ac14194d70dcebf41f0a68c89b552e94cd41cb34b9d59a23593020411  request_rss_cert-fr.go
af51b8b159bd05de8d6e3aba05dfaf70b271c6d9afc50beae56f6dc8ec5c3137  search.conf
