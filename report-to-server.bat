echo Pushing Allure-results to server and updating results
tar.exe -czf allure-results-zip.tar.gz allure-results
curl -L -X POST 'http://localhost:8010/v1/allure/addresults/user=admin/password=admin123' -F 'file=@allure-results-zip.tar.gz'
echo .
echo Server url: http://18.219.242.102:37982/index.html#