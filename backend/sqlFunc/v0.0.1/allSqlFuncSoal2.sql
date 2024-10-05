CREATE OR REPLACE FUNCTION soal2.get_all_data(datajson jsonb)
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
    result JSONB;
BEGIN

    SELECT jsonb_agg(
        jsonb_build_object(
            'id', id,
            'nameData', name_data,
            'parentId', parent_id
        )
    ) INTO result
    FROM soal2.user_data;

    -- Returning the result as a JSONB object
    RETURN result;
END;
$function$
;


CREATE OR REPLACE FUNCTION soal2.get_children_data(datajson jsonb)
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
    result JSONB;
BEGIN
    WITH RECURSIVE child_tree AS (
        SELECT id, name_data AS nameData, parent_id AS parentId
        FROM soal2.user_data
        WHERE id = (dataJson->>'uuidParent')::UUID

        UNION ALL

        SELECT u.id, u.name_data AS nameData, u.parent_id AS parentId
        FROM soal2.user_data u
        INNER JOIN child_tree ct ON u.parent_id = ct.id  -- Join with the CTE to find children
    )

   
    SELECT jsonb_agg(
        jsonb_build_object(
            'id', id,
            'nameData', ct.nameData,
           'parentId', ct.parentId
        )
    ) INTO result
    FROM child_tree ct
    where id <> (dataJson->>'uuidParent')::UUID;

    RETURN result;
END;
$function$
;
