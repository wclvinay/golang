package category
import (
  "fmt"
  "strconv"
  "encoding/json"
  "net/http"
  /*"reflect"*/
  "github.com/vanng822/go-solr/solr"
  "github.com/gorilla/mux"
)
type ResponseData struct{
	Filters []BrandName `json:"filters"`
	ProductList []SolrData `json:"productList"`
}
type SolrData struct {
	Productid float64 `json:"productid"`
	Meta_Title  string `json:"meta_title"`
	Meta_Keywords string `json:"meta_keywords"`
 	Name string `json:"name"`
	P_Brand  string `json:"p_brand"`
	P_Name  string `json:"p_name"`
	P_Pack  string `json:"p_pack"`
	Category3id float64 `json:"Category3id"`
	Root_Parent_Id float64 `json:"root_parent_id"`
	Root_Parent_Name string `json:"root_parent_name"`
	Cat_Parent_Id float64 `json:"cat_parent_id"`
	Cat_Parent_Name string `json:"cat_parent_name"`
	Cat_Name string  `json:"cat_name"`
	Status string  `json:"status"`
	Is_In_Stock float64  `json:"is_in_stock"`
	Offer_Label string  `json:"offer_label"`
	Price float64 `json:"price"`
	Sale_Price float64 `json:"sale_price"`
	Discount float64 `json:"discount"`
	Discount_Product string `json:"discount_product"`
	Is_Liked string `json:"is_liked"`
    Image string `json:"image"`
    Webqty float64 `json:"webqty"`
    Rank float64 `json:"rank"`
    Currencycode string  `json:"currencycode"`
    Product_Description string `json:"product_description"`
    Benefits string `json:"benefits"`
    Howtouse string `json:"howtouse"`
    Calories string `json:"calories"`
    Recipes string `json:"recipes"`
    Food string `json:"food"`
    Promotion_Level string `json:"promotion_level"`
}
type BrandName struct {
	Name string `json:"name"`
	Count float64  `json:"count"`
}
/*type Profile struct {
	name string
	languages []string
}*/
func CategoryData(w http.ResponseWriter, r *http.Request) {
	var main_query string
	var store_id string
	var facet_query string
	params := mux.Vars(r)
	store_id = r.Header.Get("Storeid")
	facet_query = "store_id:"+store_id
	category :=params["category"]
	si, err := solr.NewSolrInterface("http://localhost:8983/solr", "production")
	query := solr.NewQuery()
	if len(category) > 0  {
		main_query = "category:"+category
	}
	facet_query = facet_query+" AND is_visible_in_search_i:1 AND content_type:product"
	query.Q(main_query)
	fmt.Println("Filter Query:",facet_query)
	fmt.Println("Main Query:", main_query)
	query.FilterQuery(facet_query)
	query.AddFacet("brand_product_facet_s")
	query.AddFacet("hot_offer_i")
	query.AddFacetPivot("cat_name_facet_s,category_ids_i,category_boost_f")
  	s := si.Search(query)
	res, err := s.Result(nil)

	if err != nil {
		fmt.Println("Search Error %s",err)
	}
	datas := []SolrData{}
	fmt.Println("Found Records: %d",len(res.Results.Docs))
	fmt.Println("Found Records2: ",res)
	facets :=res.FacetCounts["facet_fields"]
    brand_facet := facets.(map[string]interface{})["brand_product_facet_s"].([]interface{})
	brand_array := []BrandName{}
	brand_key :="map"
	for _,brand_value :=range brand_facet{
		if brand_key =="map"{
			brand_key = brand_value.(string);
		}else {
			brand_array = append(brand_array, BrandName{
				Name : brand_key,
				Count: brand_value.(float64),
			})
			brand_key ="map"
		}

	}
	fmt.Println("First Records: %d",brand_array)
	for _, doc := range res.Results.Docs {
		fmt.Println(doc)
		var product_status string
		var product_price float64
		var product_discount float64
		var product_discount_text string
		if doc.Get("web_qty_i").(float64) >0{
			product_status = "In stock";
		}else{
			product_status = "Sold Out"
		}
		if doc.Has("actual_price_f"){
			product_price = doc.Get("actual_price_f").(float64)
			product_discount  = doc.Get("actual_price_f").(float64) - doc.Get("price_f").(float64)
		}else{
			product_price = doc.Get("price_f").(float64)
			product_discount =0
		}

		if doc.Get("discount_i").(float64) >1{
			product_discount_text = strconv.FormatFloat(doc.Get("discount_i").(float64), 'f', 4, 64)+"% OFF"
		}else{
			product_discount_text = ""
		}
		fmt.Println("here")
 		datas = append(datas, SolrData{
			Productid : doc.Get("product_id").(float64),
			Meta_Title: doc.Get("meta_title_t").(string),
			Meta_Keywords : doc.Get("meta_keywords_t").(string),
			Name :  doc.Get("name_t").(string),
			P_Brand: doc.Get("product_name_line_1_t").(string),
			P_Name: doc.Get("product_name_line_2_t").(string),
			P_Pack: doc.Get("product_name_line_2_t").(string),
			Category3id : doc.Get("category_ids_i").(float64),
			Root_Parent_Id : doc.Get("root_parent_id_i").(float64),
			Root_Parent_Name: doc.Get("root_parent_name_facet_s").(string),
			Cat_Parent_Id :  doc.Get("parent_id_i").(float64),
			Cat_Parent_Name: doc.Get("parent_name_t").(string),
			Status :product_status,
			Is_In_Stock:doc.Get("is_in_stock_i").(float64),
			Offer_Label : doc.Get("offer_label_t").(string),
			Price:product_price,
			Sale_Price : doc.Get("price_f").(float64),
			Discount : product_discount,
			Discount_Product: product_discount_text,
			Is_Liked: "true",
			Image: doc.Get("cart_image_t").(string),
			Webqty: doc.Get("web_qty_i").(float64),
			Rank: doc.Get("rank_i").(float64),
			Currencycode: "₹",
			Product_Description: doc.Get("description_t").(string),
			Benefits:doc.Get("benefits_t").(string),
		    Howtouse :doc.Get("howtouse_t").(string),
		    Calories:doc.Get("Calories_t").(string),
		    Recipes:doc.Get("recipes_t").(string),
		    Food:doc.Get("food_t").(string),
		    Promotion_Level:doc.Get("promotion_level_t").(string),
		})
 		fmt.Println("Datas",datas)
		//data = append(data,newR)
		
	}
	//fmt.Println("Records %d",datas)
	//for  i = 0; i < len(res.Results.Docs); i++ {
	//	solr_result[i] = res.Results.Docs[i]
	
	//}
  	//fmt.Println(res)
	//os.Exit()
	//profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	responseData := ResponseData{brand_array,datas}
	js,err := json.Marshal(responseData)
	fmt.Println(string(js))
	if err != nil {
   	   http.Error(w, err.Error(), http.StatusInternalServerError)
           return
        }
	w.Header().Set("Content-Type", "application/json")
  	w.Write(js)
}
