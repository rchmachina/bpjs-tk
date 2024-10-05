CREATE OR REPLACE FUNCTION soal1.get_all_data_with_pagination(datas jsonb)
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
    result JSONB;
    p_page INT := COALESCE((datas->>'page')::INT, 1); 
    p_limit INT := COALESCE((datas->>'limit')::INT, 100); 
    p_offset INT := (p_page - 1) * p_limit;  
BEGIN
    SELECT jsonb_agg(row_to_json(d.*))
    INTO result
    FROM (
        SELECT 
            id,
            nominal,
            start_date AS startDate,
            end_date AS endDate
        FROM soal1."data"
        ORDER BY created_at DESC -- Adjust the order as needed
        OFFSET p_offset LIMIT p_limit
    ) AS d;

    RETURN COALESCE(result, '[]'::jsonb); -- Return an empty JSONB array if no rows found
END;
$function$
;


CREATE OR REPLACE FUNCTION soal1.get_child_data(datajson jsonb)
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
    result JSONB;
BEGIN
    -- Selecting and aggregating the data into JSONB, sorting within the aggregate
    SELECT jsonb_agg(
        jsonb_build_object(
            'id', id,
            'startDate', start_date,
            'nominal', nominal,
            'endDate', end_date,
            'parentId', parent_id
        ) ORDER BY start_date
    ) INTO result
    FROM soal1.data_detail s
    WHERE s.parent_id = (datajson->>'parentId')::uuid;

    -- Returning the result as a JSONB object
    RETURN COALESCE(result, '[]'::jsonb); -- Return empty array if no result
END;
$function$
;


CREATE OR REPLACE FUNCTION soal1.insert_bulk_data(input_json jsonb)
 RETURNS character varying
 LANGUAGE plpgsql
AS $function$
DECLARE
    data_record JSONB;  -- Variable for each parent data record
    parent_id UUID;  -- Variable to hold the inserted parent ID
    child_record JSONB;  -- Variable for each child record
BEGIN
    -- Loop through each parent data entry in the JSON
    FOR data_record IN SELECT * FROM jsonb_array_elements(input_json->'data') LOOP
        -- RAISE NOTICE for parent data before insertion

        -- Insert into the parent table and get the inserted ID
        INSERT INTO soal1.data ( start_date, end_date, nominal)
        VALUES (
            (data_record->'parentData'->>'startDate')::DATE,  -- Cast to DATE
            (data_record->'parentData'->>'endDate')::DATE,    -- Cast to DATE
            (data_record->'parentData'->>'nominal')::NUMERIC
        )
        RETURNING id INTO parent_id;  -- Capture the inserted parent ID

        -- Raise notice for parent insertion
        RAISE NOTICE 'Inserted parent record with ID: %', parent_id;

        -- Loop through each child data entry and insert into the child table
        FOR child_record IN SELECT * FROM jsonb_array_elements(data_record->'childData') LOOP


            -- Insert into the child table
            INSERT INTO soal1.data_detail (start_date, end_date, nominal, parent_id)
            VALUES (
                (child_record->>'startDate')::DATE,  -- Cast to DATE
                (child_record->>'endDate')::DATE,    -- Cast to DATE
                (child_record->>'nominal')::NUMERIC,
                parent_id
                
                  -- Set the parent_id to the captured ID
            );

            -- Raise notice for child insertion
            RAISE NOTICE 'Inserted child record with nominal: % for parent ID: %', 
                (child_record->>'nominal'), parent_id;
        END LOOP;
    END LOOP;

    RETURN 'Success';
EXCEPTION
    WHEN OTHERS THEN
        RETURN 'Error: ' || SQLERRM;  -- Return the error message
END;
$function$
;
