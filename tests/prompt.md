## **Prompt: Generate Joint Movements JSON for an Exercise**

You are an expert in biomechanics and kinesiology. Your task is to analyze the given exercise and return a structured JSON object listing the associated joint movements based on the provided exercise name and description.

1. The associated joint movements, including their respective movement phases (concentric or eccentric).
2. The stabilizer muscles, indicating their associated joint.

### **Input Data:**
- **Exercise Name:** `Bench Press Teste`
- **Exercise Description:** `The bench press is a weight training exercise where a person presses a weight upward while lying on a bench, primarily targeting the pectoralis major, anterior deltoids, and triceps brachii. Stabilizing muscles in the back, legs, and core also assist. Typically performed with a barbell or dumbbells, it is a key lift in powerlifting alongside the squat and deadlift and the sole lift in Paralympic powerlifting. Widely used in bodybuilding and strength training, the bench press enhances upper body strength, power, and muscle development.`

### **Additional Information:**
You have access to the following reference data:
- **List of Joints:** `[{"name":"Scapula \u0026 Clavicle","id":1},{"name":"Shoulder Glenohumeral","id":2},{"name":"Elbow","id":3},{"name":"Forearm Radioulnar","id":4},{"name":"Wrist \u0026 Midcarpals","id":5},{"name":"Finger Metacarpophalangeal","id":6},{"name":"Finger Interphalangeal","id":7},{"name":"Thumb Carpometacarpal","id":8},{"name":"Thumb Metacarpophalangeal","id":9},{"name":"Thumb Interphalangeal","id":10},{"name":"Neck Atlantoccipital \u0026 Antlantoaxial","id":11},{"name":"Spine Cervical","id":12},{"name":"Spine Thoracic, Lumbar","id":13},{"name":"Hip","id":14},{"name":"Knee","id":15},{"name":"Ankle","id":16},{"name":"Foot Intertarsal","id":17},{"name":"Toe Metatarsophalangeal","id":18},{"name":"Toe Interphalangeal","id":19},{"name":"Sternoclavicular","id":20},{"name":"Costochondral","id":21},{"name":"Acromioclavicular","id":22},{"name":"Scapulothoracic","id":23},{"name":"Thoracolumbar","id":24},{"name":"Iliac crest","id":25},{"name":"Pelvis","id":26},{"name":"Sacroiliac","id":27},{"name":"Pubic crest","id":28},{"name":"Vertebrae","id":29},{"name":"Ribs","id":30},{"name":"Xiphoid process","id":31},{"name":"Temporomandibular","id":32}]`
- **List of Movements:** `[{"name":"Abduction (Protraction)","id":1},{"name":"Adduction (Retraction)","id":2},{"name":"Depression","id":3},{"name":"Elevation","id":4},{"name":"Upward Rotation (Superior Rotation)","id":5},{"name":"Downward Rotation (Inferior Rotation)","id":6},{"name":"Flexion","id":7},{"name":"Extension","id":8},{"name":"Adduction","id":9},{"name":"Abduction","id":10},{"name":"Transverse Adduction","id":11},{"name":"Transverse Flexion","id":12},{"name":"Transverse Abduction","id":13},{"name":"Transverse Extension","id":14},{"name":"Medial Rotation (Internal Rotation)","id":15},{"name":"Lateral Rotation (External Rotation)","id":16},{"name":"Pronation","id":17},{"name":"Supination","id":18},{"name":"Extension / Hyperextension","id":19},{"name":"Adduction (Ulna Deviation)","id":20},{"name":"Abduction (Radial Deviation)","id":21},{"name":"Opposition","id":22},{"name":"Lateral Flexion (Abduction)","id":23},{"name":"Reduction (Adduction)","id":24},{"name":"Rotation","id":25},{"name":"Plantar Flexion","id":26},{"name":"Dorsiflexion","id":27},{"name":"Inversion (Supination)","id":28},{"name":"Eversion (Pronation)","id":29},{"name":"Circumduction ","id":30},{"name":"Adduction (horizontal)","id":31},{"name":"Abduction (horizontal)","id":32},{"name":"Anterior tilt","id":33},{"name":"Posterior tilt","id":34},{"name":"Lateral tilt","id":35},{"name":"Nutation","id":37},{"name":"Counternutation","id":38},{"name":"Symphysis pubis","id":39},{"name":"Lateral deviation","id":40},{"name":"Hinge","id":41},{"name":"Medial Rotation of the Tibia","id":42},{"name":"Transverse Flexion","id":43}]`
- **List of Muscle Portions:** `[{"name":"Upper Major","id":1},{"name":"Medial Major","id":2},{"name":"Lower Major","id":3},{"name":"Minor","id":4},{"name":"Anterior Serratus","id":5},{"name":"Subclavio","id":6},{"name":"Latissimus dorsi","id":7},{"name":"Posterior Serratus","id":8},{"name":"Lumbar","id":9},{"name":"Upper Trapezius","id":10},{"name":"Medial Trapezius","id":11},{"name":"Lower Trapezius","id":12},{"name":"Minor Rhomboid","id":13},{"name":"Major Rhomboid","id":14},{"name":"Levator Scapulae","id":15},{"name":"Long head","id":16},{"name":"Short head","id":17},{"name":"Brachialis","id":18},{"name":"Long head","id":19},{"name":"Medial head","id":20},{"name":"Lateral head","id":21},{"name":"Pronator teres","id":22},{"name":"Flexor Carpi Radials","id":23},{"name":"Flexor Carpi Ulnaris","id":24},{"name":"Palmaris longus","id":25},{"name":"medius","id":26},{"name":"minimus","id":27},{"name":"maximus","id":28},{"name":"Pelvic floor","id":29},{"name":"Transversus Abdominis","id":30},{"name":"Multifidus","id":31},{"name":"Internal Obliques","id":32},{"name":"External Obliques","id":33},{"name":"Rectus Abdominis","id":34},{"name":"Erector spinae","id":35},{"name":"Bíceps Femoris","id":36},{"name":"Semitendinosus","id":37},{"name":"Semimembranosus","id":38},{"name":"Vastus lateralis","id":39},{"name":"Vastus medials","id":40},{"name":"Vastus intermedius","id":41},{"name":"Rectus Femoris","id":42},{"name":"Gastrocnemius","id":43},{"name":"Soleos","id":44},{"name":"Anterior","id":45},{"name":"Medium","id":46},{"name":"Lateral","id":47},{"name":"Sternocleidomastoid","id":48},{"name":"Splenius","id":49},{"name":"Tensor Fasciae Latae","id":50},{"name":"Sartorius","id":51},{"name":"Supraspinatus","id":52},{"name":"Extensor Carpi Radialis Longus","id":53},{"name":"Extensor Carpi Radialis Brevis","id":54},{"name":"Extensor Carpi Ulnaris","id":55},{"name":"Adductor longus","id":56},{"name":"Adductor brevis","id":57},{"name":"Adductor magnus","id":58},{"name":"Gracillis","id":59},{"name":"Infraspinatus","id":60},{"name":"Teres Minor","id":61},{"name":"Teres Major","id":62},{"name":"Pectineus","id":63},{"name":"Piriformis","id":64},{"name":"Quadratus Femoris","id":65},{"name":"Obturator externus","id":66},{"name":"Obturator internus","id":67},{"name":"Superior Gemellus","id":68},{"name":"Tibialis Anterior","id":69},{"name":"Subscapularis","id":70},{"name":"Iliopsoas","id":71},{"name":"Popliteus","id":72},{"name":"Posterior","id":73},{"name":"Brachioradialis","id":74},{"name":"Supinator","id":75},{"name":"Abductor magnus","id":76},{"name":"Latissimus dorsi, lower fibers","id":77},{"name":"Latissimus dorsi, upper fibers","id":78}]`

### **Output Format:**
Return a JSON array where each entry represents a joint movement involved in the exercise, formatted as follows:

```json
[
  {
    "moviment_id": <integer>,
    "moviment_name": "<string>",
    "joint_id": <integer>,
    "joint_name": "<string>",
    "muscle_id": <integer>,
    "muscle_name": "<string>",
    "role": "<concentric, eccentric or stabilizer>"
  }
]```
### Example Output:
For a Pull-Up exercise, the expected output format would be similar to:

```json
[
  {
    "moviment_id": 1,
    "moviment_name": "Flexion",
    "joint_id": 1,
    "joint_name": "Elbow",
    "muscle_id": 1,
    "muscle_name": "Biceps Brachii",
    "role": "concentric"
  },
  {
    "moviment_id": 2,
    "moviment_name": "Extension",
    "joint_id": 1,
    "joint_name": "Elbow",
    "muscle_id": 1,
    "muscle_name": "Biceps Brachii",
    "role": "eccentric"
  },
  {
    "muscle_id": 5,
    "muscle_name": "Rectus Abdominis",
    "joint_id": 3,
    "joint_name": "Spine",
    "role": "stabilizer"
  },
  ...
]
```
### Important Guidelines:
- Identify all major joint movements involved in the exercise.
- Map each joint movement to its corresponding muscle portion.
- Use the provided reference lists to ensure consistency in naming and IDs.
- Output only valid movements relevant to the exercise mechanics.
- The response should be in pure JSON format, without additional explanations.

