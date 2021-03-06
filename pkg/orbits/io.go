package orbits

import (
   "fmt"
	"os"
   "strconv"
	"io/ioutil"
	
   "github.com/asimazbunzel/go-orbits/pkg/io"

	"gopkg.in/yaml.v3"
)


// read information on binary and kicks from YAML file
func (b *Binary) parseYAML (filename string) error {

   // read YAML data file into bytes 
   data, err := ioutil.ReadFile(filename)
   if err != nil {
      io.LogError("ORBITS - orbits.go - ParseYAML", "problem reading YAML file")
   }
   
   return yaml.Unmarshal(data, b)
}


// save kick info to file
func (b *Binary) SaveKicks (filename string) {

   if b.LogLevel != "none"{
      io.LogInfo("ORBITS - orbits.go - SaveKicks", "saving kicks information")
   }

   // create file
   f, err := os.Create(filename)
   if err != nil {
      io.LogError("error writing to file", "open file")
   }

   // remember to close the file
   defer f.Close()

   // header
   column_names := [4]string{"id", "w", "theta", " phi"}
   str := fmt.Sprintf("%20s", column_names[0]) 
   str += fmt.Sprintf("%20s", column_names[1])
   str += fmt.Sprintf("%20s", column_names[2])
   str += fmt.Sprintf("%20s\n", column_names[3])
   _, err = f.WriteString(str)
   if err != nil {
      io.LogError("ORBITS - orbits.go - SaveKicks", "error writing header to file")
   }

   // write rows of different natal kicks
   for k, w := range b.W {
      str := fmt.Sprintf("%20s", strconv.Itoa(k))
      str += fmt.Sprintf("%20s", strconv.FormatFloat(w, 'E', 5, 64))
      str += fmt.Sprintf("%20s",strconv.FormatFloat(b.Theta[k], 'E', 5, 64))
      str += fmt.Sprintf("%20s\n",strconv.FormatFloat(b.Phi[k], 'E', 5, 64))
      _, err := f.WriteString(str)
      if err != nil {
         io.LogError("ORBITS - orbits.go - SaveKicks", "error writing info to file")
      }
   }

}


// save orbits info to file
func (b *Binary) SaveBoundedOrbits (filename string) {

   if b.LogLevel != "none"{
      io.LogInfo("ORBITS - orbits.go - SaveBoundedOrbits", "saving orbits information")
   }

   // create file
   f, err := os.Create(filename)
   if err != nil {
      io.LogError("error writing to file", "open file")
   }

   // remember to close the file
   defer f.Close()

   // header
   column_names := [7]string{"id", "w", "theta", "phi", "period", "separation", "eccentricity"}
   str := fmt.Sprintf("%20s", column_names[0]) 
   str += fmt.Sprintf("%20s", column_names[1])
   str += fmt.Sprintf("%20s", column_names[2])
   str += fmt.Sprintf("%20s", column_names[3])
   str += fmt.Sprintf("%20s", column_names[4]) 
   str += fmt.Sprintf("%20s", column_names[5])
   str += fmt.Sprintf("%20s\n", column_names[6])
   _, err = f.WriteString(str)
   if err != nil {
      io.LogError("ORBITS - orbits.go - SaveBoundedOrbits", "error writing header to file")
   }

   // write rows of different natal kicks
   for k, kb := range b.IndexBounded {
      str := fmt.Sprintf("%20s", strconv.Itoa(kb))
      str += fmt.Sprintf("%20s", strconv.FormatFloat(b.WBounded[k], 'E', 5, 64))
      str += fmt.Sprintf("%20s", strconv.FormatFloat(b.ThetaBounded[k], 'E', 5, 64))
      str += fmt.Sprintf("%20s", strconv.FormatFloat(b.PhiBounded[k], 'E', 5, 64))
      str += fmt.Sprintf("%20s", strconv.FormatFloat(b.PeriodBounded[k], 'E', 5, 64))
      str += fmt.Sprintf("%20s",  strconv.FormatFloat(b.SeparationBounded[k], 'E', 5, 64))
      str += fmt.Sprintf("%20s\n",  strconv.FormatFloat(b.EccentricityBounded[k], 'E', 5, 64))
      _, err := f.WriteString(str)
      if err != nil {
         io.LogError("ORBITS - orbits.go - SaveBoundedOrbits", "error writing info to file")
      }
   }

}


// save grid of binaries bounded after kick
func (b *Binary) SaveGridOrbits (filename string) {

   if b.LogLevel != "none"{
      io.LogInfo("ORBITS - orbits.go - SaveGridOrbits", "saving grid of orbits information")
   }

   // create file
   f, err := os.Create(filename)
   if err != nil {
      io.LogError("error writing to file", "open file")
   }

   // remember to close the file
   defer f.Close()

   // header
   column_names := [5]string{"id", "period", "separation", "eccentricity", "probability"}
   str := fmt.Sprintf("%20s", column_names[0]) 
   str += fmt.Sprintf("%20s", column_names[1])
   str += fmt.Sprintf("%20s", column_names[2])
   str += fmt.Sprintf("%20s", column_names[3])
   str += fmt.Sprintf("%20s\n", column_names[4]) 
   _, err = f.WriteString(str)
   if err != nil {
      io.LogError("ORBITS - orbits.go - SaveGridOrbits", "error writing header to file")
   }

   // write info to file
   for k, _ := range b.PeriodGrid {
      str := fmt.Sprintf("%20s", strconv.Itoa(k))
      str += fmt.Sprintf("%20s", strconv.FormatFloat(b.PeriodGrid[k], 'E', 5, 64))
      str += fmt.Sprintf("%20s", strconv.FormatFloat(b.SeparationGrid[k], 'E', 5, 64))
      str += fmt.Sprintf("%20s", strconv.FormatFloat(b.EccentricityGrid[k], 'E', 5, 64))
      str += fmt.Sprintf("%20s\n", strconv.FormatFloat(b.ProbabilityGrid[k], 'E', 5, 64))
      _, err := f.WriteString(str)
      if err != nil {
         io.LogError("ORBITS - orbits.go - SaveGridOrbits", "error writing info to file")
      }
   }

}
