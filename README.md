# SellerApp

Below I have listed all the details to run this project.

- Go to the root directory - that is in the sellerApp directory and there type the below command to build and run the containers

  - ```dockerfile
    docker-compose up --build
    ```

- Now all the three containers will be running, so you just have to open the postman and send a POST request on 

  ```
  http://localhost:8000/crawl
  ```

   with body input in JSON format as

  ```json
  {"url":"https://www.amazon.com/dp/B07SH57BCY/ref=sspa_dk_detail_0?psc=1&pd_rd_i=B07SH57BCY&pd_rd_w=rv6Ca&pf_rd_p=5d846283-ed3e-4512-a744-a30f97c5d738&pd_rd_wg=8k0DC&pf_rd_r=KC781B3KTG8VCBG19X5B&pd_rd_r=7164a4eb-2476-4fd0-ae10-507244a0948c&spLa=ZW5jcnlwdGVkUXVhbGlmaWVyPUExUTA5WENQVjQ4WlJUJmVuY3J5cHRlZElkPUEwMDQwMTg1MTdMNFFMUENJSk0xUCZlbmNyeXB0ZWRBZElkPUEwMDM4MDUxMUxBVlA1UllCRUFYQSZ3aWRnZXROYW1lPXNwX2RldGFpbF90aGVtYXRpYyZhY3Rpb249Y2xpY2tSZWRpcmVjdCZkb05vdExvZ0NsaWNrPXRydWU="}
  ```

- The request will be sent to the fetcher service and after crawling data the fetcher service further calls to the data writer service which writes the product details into mongodb database.

- The Ouput format will be

  ```json
  {
      "url": "https://www.amazon.com/dp/B07SH57BCY/ref=sspa_dk_detail_0?psc=1&pd_rd_i=B07SH57BCY&pd_rd_w=rv6Ca&pf_rd_p=5d846283-ed3e-4512-a744-a30f97c5d738&pd_rd_wg=8k0DC&pf_rd_r=KC781B3KTG8VCBG19X5B&pd_rd_r=7164a4eb-2476-4fd0-ae10-507244a0948c&spLa=ZW5jcnlwdGVkUXVhbGlmaWVyPUExUTA5WENQVjQ4WlJUJmVuY3J5cHRlZElkPUEwMDQwMTg1MTdMNFFMUENJSk0xUCZlbmNyeXB0ZWRBZElkPUEwMDM4MDUxMUxBVlA1UllCRUFYQSZ3aWRnZXROYW1lPXNwX2RldGFpbF90aGVtYXRpYyZhY3Rpb249Y2xpY2tSZWRpcmVjdCZkb05vdExvZ0NsaWNrPXRydWU=",
      "product": {
          "imageUrl": "https://images-na.ssl-images-amazon.com/images/I/81KfqjC8LAL._AC_SL1500_.jpg",
          "name": "NINETYGO Carry on Luggage 22x14x9 with Spinner Wheels, Hardside Carry on Suitcase with Front Pocket Lock Cover, Super Convenient & Lightweight for Business Travel (20\")",
          "description": "Make sure this fits\nby entering your model number.\n\n\nLightweight & Durable Hardshell: This carry on luggage is built from Covestro (former Germany BAYER MaterialScience). The 100% Makrolon Polycarbonate of the hardshell luggage is incredibly strong, scratch-resistant, lighter and more durable than aluminum and ABS. Package Dimensions: 15.0\" X 9.5\" X 22.5\", Capacity: 36 L, Weight: 6.4 lbs.\n\n\nUltra-silent & Multi-directional Spinner Dual Wheels: This 4-wheel spinner of carry on luggage 22x14x9 features wear-resistant, whisper-silent. Spinner wheels with TPE vibration dampen for increased lifespan and smoothness. The quiet smooth wheels of this spinner luggage turn 360 degrees making it easier to navigate around the terminal or narrow aisles.\n\n\nRetractable Handle with Flexible Heights: The 4 wheeled carry on luggage has a rigid 4-level telescoping push-button handle which locks firmly in place to give you ultimate control and comfort (shaking distance of the handle ≤ 0.6\").\n\n\nTSA Compatible Combination Zipper Lock: Acts to prevent theft, NO keys needed! Have peace of mind knowing your case is secured and your valuable items are fully protected in rolling luggage cover. No need to prepare in advance or worry about a damaged lock when a security agent inspects your hardshell luggage.\n\n\n5-Year Warranty & Satisfaction Guarantee: Our carry on suitcases have been tested to last the distance. If you have ANY concerns about your hardside luggage, please contact us and we will resolve the issue. Travel with our carry on suitcases and make your next journey a smooth ride!",
          "price": "$159.99",
          "totalReviews": 295
      }
  }
  ```

  

Thank you.