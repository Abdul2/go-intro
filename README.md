### simple RESTfull example


### intro


simple go programme to be used as part of internal training.


### to generate test data use :


```[
  
 '{{repeat(5, 200)}}',

	{	

      "personid":'{{guid()}}',

      "object":'{{lorem()}}',

      "location":'{{city()}}',

		"event":

		{

          "date":'{{date()}}',

          "eventtype":'{{lorem()}}'

		}
	}
	
]'''

 in http://www.json-generator.com/