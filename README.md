
### **DailySync API**  
**Your Daily Source for Weather, Surf, Tides, Celebrations, and Bitcoin Prices**

DailySync API is a powerful Golang-based API that provides essential daily information in one place. Whether you're planning your day, hitting the surf, or staying updated on crypto trends, DailySync has you covered.  

#### **Features**  
- **Weather Conditions**: Get real-time weather updates for your location.  
- **Surf Conditions**: Check the latest surf report for your favorite spots.  
- **Tide Status**: Monitor high and low tide percentages in real time.  
- **Today’s Celebrations**: Discover whose name day or celebration it is today.  
- **Bitcoin Price**: Stay informed with the latest BTC price.  

#### **Why DailySync?**  
- **All-in-One Solution**: Access multiple daily insights through a single API.  
- **Secure Access**: Protected with JWT-based Bearer Token authentication.  
- **Easy Integration**: Built in Golang for high performance and scalability.  

DailySync API is perfect for developers, surfers, weather enthusiasts, and crypto traders who need quick, reliable daily updates.  

---

### **TO DO**

- [ ] Add **Smart Caching System**  
  - Implement an in-memory cache using `go-cache` to store API responses and reduce redundant external API calls.  
  - Set cache expiration times tailored to each data type (e.g., weather, tides, Bitcoin price).  
  - Ensure cache invalidation when data becomes stale or outdated.  

- [ ] Add **Health Check Endpoint**  
  - Create a `/health` endpoint to monitor the status of the API and its dependencies.  


---
