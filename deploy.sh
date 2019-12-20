echo Deploying from $0

echo Installing all necessary packages
cd ../BotSecrets
git pull
go install
cd ../ScriptureBot
git pull
go install
cd $0
git pull
go install

echo Deploying app
gcloud app deploy --quiet