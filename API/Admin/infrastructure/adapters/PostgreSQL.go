package adapters

import (
	"GoAir-Admin/API/Admin/domain/entities"
	"GoAir-Admin/API/core"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostgreSQL struct {
	conn *core.ConnPostgreSQL
}

func NewPostgreSQL() *PostgreSQL {
	conn := core.GetDBPool()

	if conn.Err != "" {
		fmt.Println("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &PostgreSQL{conn: conn}
}

func (postgres *PostgreSQL) GetAdmin() entities.Admin {
	query := "SELECT * FROM admin"
	var admin entities.Admin

	rows, err := postgres.conn.FetchRows(query)

	if err != nil {
		fmt.Errorf("error al ejecutar la consulta: %w", err)
		return entities.Admin{}
	}

	defer rows.Close()

	if !rows.Next() {
        fmt.Println("No se pudieron obtener los datos.")
        return entities.Admin{}
    }

	if err := rows.Scan(&admin.Password, &admin.Email); err != nil {
		fmt.Errorf("error al escanear el admin: %w", err)
        return entities.Admin{}
	}
	fmt.Print(admin.Password)

	return admin
}

func (postgres *PostgreSQL) SearchUser(last_name string) entities.User {
	query := "SELECT * FROM users WHERE last_name LIKE '%' || $1 || '%'"
	var user entities.User
	
	rows, err := postgres.conn.FetchRows(query, last_name)

	if err != nil {
		fmt.Errorf("error al ejecutar la consulta: %w", err)
		return entities.User{}
	}

	defer rows.Close()

	if !rows.Next() {
        fmt.Println("No se pudieron obtener los datos.")
        return entities.User{}
    }

	if err := rows.Scan(&user.Id_user, &user.First_name, &user.Last_name, &user.Email, &user.Password); err != nil {
		fmt.Errorf("error al escanear el usuario: %w", err)
        return entities.User{}
    }
	fmt.Print(user)
	return user
}

func (postgres *PostgreSQL) CreatePlace(name string, id_user int, id_application int) (uint, error){
	query := "INSERT INTO places (name, id_user) VALUES ($1, $2) RETURNING id_place"

	var id uint
	err := postgres.conn.DB.QueryRow(query, name, id_user).Scan(&id)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta 1: %v", err)
		return 0, err
	}

	if err = postgres.CreateId(int(id)); err != nil {
		fmt.Println("Error: %v", err)
		return 0, err
	} 

	if err = changeStatusApplication(postgres, id_application); err != nil {
		fmt.Println("Error al cambiar el estado: %v", err)
		return 0, err
	}


	return id, nil
}

func changeStatusApplication(p *PostgreSQL, id_application int) error {
    query := "UPDATE applications SET status_application = 'pending' WHERE id_application = $1"
    
    result, err := p.conn.DB.Exec(query, id_application)
    if err != nil {
        return fmt.Errorf("error actualizando estado: %w", err)
    }

    // Verificar si realmente se actualizó algún registro
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("ningún registro actualizado (id_application %d no existe)", id_application)
    }

    return nil
}

func (postgres *PostgreSQL) CreateId(id_place int) (error) {
	id_mq135a := primitive.NewObjectID().Hex() 
	id_mq135b := primitive.NewObjectID().Hex()  
	id_dh11a := primitive.NewObjectID().Hex()  
	id_dh11b := primitive.NewObjectID().Hex()
	id_deviceA := primitive.NewObjectID().Hex() 
	id_deviceB := primitive.NewObjectID().Hex() 

	query := "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES ($1,$2,$3,$4)"
	_, err := postgres.conn.DB.Query(query, id_mq135a, id_place, "air_quality", "MQ135")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaA: %v", err)
		return err
	}

	query = "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES ($1,$2,$3,$4)"
	_, err = postgres.conn.DB.Query(query, id_mq135b, id_place, "air_quality", "MQ135")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaB: %v", err)
		return err
	}

	query = "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES ($1,$2,$3,$4)"
	_, err = postgres.conn.DB.Query(query, id_dh11a, id_place, "temperature", "DH11")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaC: %v", err)
		return err
	}

	query = "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES ($1,$2,$3,$4)"
	_, err = postgres.conn.DB.Query(query, id_dh11b, id_place, "temperature", "DH11")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaD: %v", err)
		return err
	}

	query = "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES ($1,$2,$3,$4)"
	_, err = postgres.conn.DB.Query(query, id_dh11a, id_place, "humidity", "DH11")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaE: %v", err)
		return err
	}

	query = "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES ($1,$2,$3,$4)"
	_, err = postgres.conn.DB.Query(query, id_dh11b, id_place, "humidity", "DH11")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaF: %v", err)
		return err
	}

	query = "INSERT INTO devices (id_device, id_place) VALUES ($1,$2)"
	_, err = postgres.conn.DB.Query(query, id_deviceA, id_place)
	if err != nil {
		fmt.Println("Error al ejecutar la consultaF: %v", err)
		return err
	}

	query = "INSERT INTO devices (id_device, id_place) VALUES ($1,$2)"
	_, err = postgres.conn.DB.Query(query, id_deviceB, id_place)
	if err != nil {
		fmt.Println("Error al ejecutar la consultaF: %v", err)
		return err
	}

	return nil
}

func (postgres *PostgreSQL) GetIds(id_place int) ([]entities.Sensor, []entities.Device) {
	query := "SELECT * FROM sensors WHERE id_place = $1"
	var sensors []entities.Sensor

	rows, err := postgres.conn.DB.Query(query, id_place)

	if err != nil {
        fmt.Println("No se pudieron obtener los datos.", err)
        return []entities.Sensor{}, []entities.Device{}
    }

	defer rows.Close()

	for rows.Next() {
		var s entities.Sensor
		rows.Scan(&s.Id_sensor, &s.Id_place, &s.Sensor_type, &s.Model, &s.Installation_date)

		sensors = append(sensors, s)
	}
	devices :=  getDevices(postgres, id_place)
	return sensors, devices
}

func getDevices(postgres *PostgreSQL, id_place int) []entities.Device {
	query := "SELECT * FROM devices WHERE id_place = $1"
	var devices []entities.Device

	rows, err := postgres.conn.DB.Query(query, id_place)

	if err != nil {
        fmt.Println("No se pudieron obtener los datos.", err)
        return []entities.Device{}
    }

	defer rows.Close()

	for rows.Next() {
		var d entities.Device
		rows.Scan(&d.Id_device, &d.Id_place)

		devices = append(devices, d)
	}
	
	return devices
}

func (postgres *PostgreSQL) GetPlaces(id_user int) []entities.Place {
	query := "SELECT * FROM places WHERE id_user = $1"
	fmt.Println(id_user)
	var places []entities.Place

	rows, err := postgres.conn.DB.Query(query, id_user)

	if err != nil {
        fmt.Println("No se pudieron obtener los datos.", err)
        return []entities.Place{}
    }

	defer rows.Close()

	for rows.Next() {
		var p entities.Place
		
		// Escanear los valores de la fila
		err := rows.Scan(&p.Id_place, &p.Id_user, &p.Name, &p.Create_at)
		if err != nil {
			// Manejar error al escanear la fila
			fmt.Println("Error al escanear la fila:", err)
			return []entities.Place{}
		}
		places = append(places, p)
	}

	// Verifica errores después de iterar
    if err = rows.Err(); err != nil {
        fmt.Println("Error al recorrer las filas:", err)
        return nil
    }

	return places
}

func (postgres *PostgreSQL) DeletePlace(id_place int) (uint, error) {
	query := "DELETE FROM places WHERE id_place = $1"
	
	_, err := postgres.conn.DB.Exec(query, id_place)


	if err != nil {
		fmt.Println("Error al ejecutar la consultaF: %v", err)
		return 0, err
	}
	
	return uint(1), nil
}

func (postgres *PostgreSQL) GetApplicationByUser(id_user int) []entities.Application {
	query := `SELECT * FROM applications WHERE id_user = $1 AND status_application != 'complete'`

	var applications []entities.Application

	rows, err := postgres.conn.DB.Query(query, id_user)

	if err != nil {
        fmt.Println("No se pudieron obtener los datos.", err)
        return []entities.Application{}
    }

	defer rows.Close()

	for rows.Next() {
		var a entities.Application
		
		// Escanear los valores de la fila
		err := rows.Scan(&a.Id_application, &a.Status_application, &a.Id_user)
		if err != nil {
			// Manejar error al escanear la fila
			fmt.Println("Error al escanear la fila:", err)
			return []entities.Application{}
		}
		applications = append(applications, a)
	}

	// Verifica errores después de iterar
    if err = rows.Err(); err != nil {
        fmt.Println("Error al recorrer las filas:", err)
        return nil
    }

	return applications
}

func (postgres *PostgreSQL) GetAllApplications() ([]entities.AllApplications) {
	query := `SELECT a.id_application, u.first_name, u.last_name, a.status_application, a.id_user
			  FROM applications a
			  INNER JOIN users u 
			  ON a.id_user = u.id_user
			  LIMIT 100`

	var applications []entities.AllApplications

	rows, err := postgres.conn.DB.Query(query)
		  
	if err != nil {
		fmt.Println("No se pudieron obtener los datos.", err)
		return []entities.AllApplications{}
	}
		  
	defer rows.Close()
		  
	for rows.Next() {
		var a entities.AllApplications
				  
		// Escanear los valores de la fila
		err := rows.Scan(&a.Id_application, &a.First_name, &a.Last_name, &a.Status_application, &a.Id_user)
		if err != nil {
			// Manejar error al escanear la fila
			fmt.Println("Error al escanear la fila:", err)
			return []entities.AllApplications{}
			}
			applications = append(applications, a)
		}
		  
		// Verifica errores después de iterar
		if err = rows.Err(); err != nil {
			fmt.Println("Error al recorrer las filas:", err)
			return nil
	}
		  
	return applications
}

func (postgres *PostgreSQL) ConfirmInstallation(id_application int) (int, error) {
	query := "SELECT id_user FROM applications WHERE id_application = $1"	

	var id_user int 

	rows, err := postgres.conn.FetchRows(query, id_application)
	if err != nil {
		fmt.Errorf("error al ejecutar la consulta: %w", err)
		return 0, err
	}
	
	if !rows.Next() {
        fmt.Println("No se pudieron obtener los datos.")
        return 0, err
    }

	if err := rows.Scan(&id_user); err != nil {
		fmt.Errorf("error al escanear el usuario: %w", err)
        return 0, err
	}
	
	query = "DELETE FROM applications WHERE id_application = $1"
	
	result, err := postgres.conn.DB.Exec(query, id_application)

	if err != nil {
		fmt.Println("Error al ejecutar la consultaF: %v", err)
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return 0, fmt.Errorf("ningún registro actualizado (id_application %d no existe)", id_application)
    }

	
	return id_user, nil
}